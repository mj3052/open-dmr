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
		if vehicle.Plate == "" {
			continue
		}

		buffer = append(buffer, vehicle)
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
