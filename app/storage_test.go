package main

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestStorePlates(t *testing.T) {
	os.RemoveAll("vehicles-test.db")

	db, err := gorm.Open(sqlite.Open("vehicles-test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	vdb = &VehicleDB{db}
	vdb.Migrate()

	v, err := vdb.VehicleLookup("XP22655")

	if err == nil {
		t.Error("Found plate before loading data, test is invalid")
	}

	vdb.LoadData("test_data/snippet.xml.zip")

	v, err = vdb.VehicleLookup("XP22655")

	if err != nil {
		t.Error("Error while looking up plate:", err, v)
	}

	if v.Plate != "XP22655" {
		t.Errorf("Test plate not correct??")
	}
}
