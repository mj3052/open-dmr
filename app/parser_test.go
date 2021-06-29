package main

import (
	"testing"
)

var testData []Vehicle

func getTestData() []Vehicle {
	if testData != nil {
		return testData
	}

	c := ParseDataFile("test_data/snippet.xml.zip", "Statistik", true)
	buffer := make([]Vehicle, 0, 10)
	for vehicle := range c {
		buffer = append(buffer, vehicle)
	}
	return buffer
}

func TestParseXML(t *testing.T) {
	c := ParseDataFile("test_data/snippet.xml", "Statistik", false)

	buffer := make([]Vehicle, 0, 10)

	for vehicle := range c {
		buffer = append(buffer, vehicle)
	}

	if len(buffer) != 8 {
		t.Errorf("Did not parse all items, parsed %d of 8", len(buffer))
	}

}

func TestParseZip(t *testing.T) {
	buffer := getTestData()
	if len(buffer) != 8 {
		t.Errorf("Did not parse all items, parsed %d of 8", len(buffer))
	}
}
