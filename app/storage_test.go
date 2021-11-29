package main

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestStorePlates(t *testing.T) {
	os.Remove("test-db.db")

	db, err := gorm.Open(sqlite.Open("test-db.db"), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Vehicle{})

	vdb = &VehicleDB{db}

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

	os.Remove("test-db.db")
}
