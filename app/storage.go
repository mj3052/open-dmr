package main

import (
	"fmt"

	"gorm.io/gorm"
)

type VehicleDB struct {
	db *gorm.DB
}

type VehicleNotFound struct{}

func (m *VehicleNotFound) Error() string {
	return "Not found"
}

func (vdb *VehicleDB) VehicleLookup(plate string) (Vehicle, error) {
	var vehicle Vehicle
	vdb.db.First(&vehicle, "plate = ?", plate)
	if vehicle.Plate != "" {
		return vehicle, nil
	}
	return vehicle, &VehicleNotFound{}
}

func (vdb *VehicleDB) LoadData(filename string) {
	p := ParseDataFile(filename, "Statistik", true)

	// Initialize buffer
	batchSize := 10000
	buffer := make([]Vehicle, 0, batchSize)

	i := 0
	for vehicle := range p {

		// Skip empty
		if vehicle.Plate == "" || vehicle.BaseInfo.Status == "Afmeldt" {
			continue
		}

		buffer = append(buffer, vehicle)

		if len(buffer) == batchSize {
			vdb.db.CreateInBatches(buffer, 1000)
			buffer = nil
			buffer = make([]Vehicle, 0, batchSize)
		}

		i++

		if i%1000 == 0 {
			fmt.Printf("%d items inserted\n", i)
		}

	}
	// Flush buffer on end
	if len(buffer) > 0 {
		vdb.db.CreateInBatches(buffer, 1000)
		buffer = nil
	}
}
