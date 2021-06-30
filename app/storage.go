package main

import (
	"encoding/json"
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

// VehicleDB contains db insert+lookup functions
type VehicleDB struct {
	db *badger.DB
}

// Insert a batch of vehicles into the DB
func (vdb *VehicleDB) processBatch(batch []Vehicle) {
	err := vdb.db.Update(func(txn *badger.Txn) error {
		for _, v := range batch {
			b, _ := json.Marshal(v)
			txn.Set([]byte("plates/"+v.Plate), b)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

// LoadData loads the given DMV zip into the DB
func (vdb *VehicleDB) LoadData(filename string) {
	p := ParseDataFile(filename, "Statistik", true)

	// Initialize buffer
	batchSize := 5000
	buffer := make([]Vehicle, 0, batchSize)

	i := 0
	for vehicle := range p {

		// Skip empty
		if vehicle.Plate == "" || vehicle.BaseInfo.Status == "Afmeldt" {
			continue
		}

		buffer = append(buffer, vehicle)
		i++

		// Flush buffer every time it's full
		if i%batchSize == 0 {
			// vdb.processBatch(buffer)
			buffer = buffer[:0] // Slice to 0 to reuse memory for next batch
			fmt.Printf("%d items inserted\n", i)
		}

		// Break early right now
		if i > 500000 {
			break
		}

	}

	// Make sure to use last bit of buffer
	vdb.processBatch(buffer)
}

// VehicleLookup looks for a plate
func (vdb *VehicleDB) VehicleLookup(plate string) (Vehicle, error) {
	var vehicle Vehicle

	err := vdb.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("plates/" + plate))

		if err != nil {
			return err
		}

		err = item.Value(func(v []byte) error {
			json.Unmarshal(v, &vehicle)
			return err
		})

		return err
	})

	return vehicle, err
}
