package main

import (
	"os"
	"testing"

	"github.com/estraier/tkrzw-go"
)

func TestStorePlates(t *testing.T) {
	os.Remove("test-db.tkh")

	dbm := tkrzw.NewDBM()
	dbm.Open("test-db.tkh", true,
		tkrzw.ParseParams("truncate=true,num_buckets=10000"))

	vdb = &VehicleDB{dbm}

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
