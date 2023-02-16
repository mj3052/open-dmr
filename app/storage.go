package main

import (
	"fmt"

	"gorm.io/gorm"
)

type VehicleDB struct {
	// db *tkrzw.DBM
	db *gorm.DB
}

type VehicleNotFound struct{}

func (m *VehicleNotFound) Error() string {
	return "Not found"

}

func (vdb *VehicleDB) VehicleLookup(plate string) (*Vehicle, error) {
	var vehicle Vehicle

	vdb.db.Take(&vehicle, "plate = ?", plate)

	// err := vdb.db.View(func(txn *badger.Txn) error {
	// 	// Your code here…
	// 	item, err := txn.Get([]byte(strings.ToUpper(plate)))

	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = item.Value(func(val []byte) error {
	// 		json.Unmarshal([]byte(val), &vehicle)
	// 		return nil
	// 	})
	// 	return nil
	// })

	if vehicle.Plate != "" {
		return &vehicle, nil
	}

	return &vehicle, &VehicleNotFound{}
}

func (vdb *VehicleDB) Migrate() {
	vdb.db.AutoMigrate(&Vehicle{})
}

func (vdb *VehicleDB) LoadData(filename string) {
	p := ParseDataFile(filename, "Statistik", true)

	i := 0

	bufferSize := 5000
	upsertBatchSize := 1000

	var buffer = make([]Vehicle, 0, bufferSize)

	for vehicle := range p {

		// Skip empty
		if vehicle.Plate == "" || vehicle.BaseInfo.Status == "Afmeldt" {
			continue
		}

		buffer = append(buffer, vehicle)

		// v, err := json.Marshal(vehicle)

		// if err != nil {
		// 	log.Error(err)
		// 	fmt.Printf("Error: %+v\n", err)
		// 	continue
		// }

		// err = wb.Set([]byte(strings.ToUpper(vehicle.Plate)), v)
		i++

		if len(buffer) == bufferSize {
			vdb.db.CreateInBatches(buffer, upsertBatchSize)
			buffer = buffer[:0]
			fmt.Printf("%d items inserted\n", i)
		}
	}

	if len(buffer) > 0 {
		vdb.db.CreateInBatches(buffer, upsertBatchSize)
	}

}

// func (vdb *VehicleDB) VehicleLookup(plate string) (Vehicle, error) {
// 	var vehicle Vehicle
// 	vdb.db.First(&vehicle, "plate = ?", plate)
// 	if vehicle.Plate != "" {
// 		return vehicle, nil
// 	}
// 	return vehicle, &VehicleNotFound{}
// }

// func (vdb *VehicleDB) LoadData(filename string) {
// 	p := ParseDataFile(filename, "Statistik", true)

// 	// Initialize buffer
// 	batchSize := 10000
// 	buffer := make([]Vehicle, 0, batchSize)

// 	i := 0
// 	for vehicle := range p {

// 		// Skip empty
// 		if vehicle.Plate == "" || vehicle.BaseInfo.Status == "Afmeldt" {
// 			continue
// 		}

// 		buffer = append(buffer, vehicle)

// 		if len(buffer) == batchSize {
// 			vdb.db.CreateInBatches(buffer, 1000)
// 			buffer = nil
// 			buffer = make([]Vehicle, 0, batchSize)
// 		}

// 		i++

// 		if i%1000 == 0 {
// 			fmt.Printf("%d items inserted\n", i)
// 		}

// 	}
// 	// Flush buffer on end
// 	if len(buffer) > 0 {
// 		vdb.db.CreateInBatches(buffer, 1000)
// 		buffer = nil
// 	}
// }
