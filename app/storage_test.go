package main

import (
	"log"
	"os"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"
)

func TestStorePlates(t *testing.T) {
	os.RemoveAll("badger-test.db")

	db, err := badger.Open(badger.DefaultOptions("./badger-test.db").WithCompression(options.ZSTD))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	os.Remove("test-db.tkh")
}
