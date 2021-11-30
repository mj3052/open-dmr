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
	// batchSize := 5000
	// buffer := make([]Vehicle, 0, batchSize)

	i := 0
	for vehicle := range p {

		// Skip empty
		if vehicle.Plate == "" || vehicle.BaseInfo.Status == "Afmeldt" {
			continue
		}
		vdb.db.Create(vehicle)

		// buffer = append(buffer, vehicle)
		i++

		if i%1000 == 0 {
			fmt.Printf("%d items inserted\n", i)
		}

		// // Flush buffer every time it's full
		// if i%batchSize == 0 {
		// 	// vdb.processBatch(buffer)
		// 	vdb.db.CreateInBatches(buffer, 1000)

		// 	buffer = buffer[:0] // Slice to 0 to reuse memory for next batch
		// 	fmt.Printf("%d items inserted\n", i)
		// }
	}
	// Flush buffer on end
	// vdb.db.CreateInBatches(buffer, 1000)
}
