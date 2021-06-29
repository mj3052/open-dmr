package main

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"log"
	"os"
)

func loadZippedXML(filePath string) io.ReadCloser {
	zf, err := zip.OpenReader(filePath)

	if err != nil {
		panic(err)
	}

	for _, file := range zf.File {

		if file.Name[0:1] == "_" || file.Name[len(file.Name)-3:] != "xml" {
			continue
		}

		f, err := file.Open()

		if err != nil {
			panic(err)
		}

		return f
	}

	return nil
}

func loadXML(filePath string) io.ReadCloser {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil
	}
	return file
}

// ParseDataFile parses a DMV data file at the given path
func ParseDataFile(filePath string, containerElement string, isZipped bool) chan Vehicle {
	// Open reader
	var file io.ReadCloser
	if isZipped {
		file = loadZippedXML(filePath)
	} else {
		file = loadXML(filePath)
	}

	// Ready decoder
	decoder := xml.NewDecoder(file)

	// Prepare channel
	c := make(chan Vehicle)

	// Stream the xml
	go func() {
		for {
			// Read tokens from the XML document in a stream.
			t, _ := decoder.Token()

			// Break if we reach the end
			if t == nil {
				break
			}

			// If token is not a StartElement or not the element we are looking for
			element, ok := t.(xml.StartElement)
			if !ok || element.Name.Local != containerElement {
				continue
			}

			// Decode the element
			var vehicle Vehicle
			decoder.DecodeElement(&vehicle, &element)

			c <- vehicle // Emit vehicle to channel
		}
		close(c)
		file.Close() // Make sure to close file and channel
	}()

	return c
}
