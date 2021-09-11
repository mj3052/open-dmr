package main

import (
	"log"
	"testing"

	badger "github.com/dgraph-io/badger/v3"
)

func TestStorePlates(t *testing.T) {
	db, err := badger.Open(badger.DefaultOptions("./test-database"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.DropAll()

	vdb = &VehicleDB{db}

	v, err := vdb.VehicleLookup("XP22655")

	if err == nil {
		t.Error("Found plate before loading data, test is invalid")
	}

	vdb.LoadData("test_data/snippet.xml.zip")

	v, err = vdb.VehicleLookup("XP22655")

	if err != nil {
		t.Error("Error while looking up plate:", err)
	}

	if v.Plate != "XP22655" {
		t.Errorf("Test plate not correct??")
	}

	db.DropAll()
}
