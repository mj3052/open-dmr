package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var vdb *VehicleDB

type jsonError struct {
	Message string `json:"message"`
}

func plateLookup(c echo.Context) error {
	plate := c.Param("plate")
	vehicle, err := vdb.VehicleLookup(plate)

	if err != nil {
		fmt.Printf("%+v\n", err)
		return c.JSON(404, jsonError{"Plate Not found :)"})
	}

	return c.JSONPretty(200, vehicle, "    ")
}

func server(host string) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/:plate", plateLookup)

	// Start server
	e.Logger.Fatal(e.Start(host))
}

func main() {
	var host = flag.String("host", "0.0.0.0", "Load data to DB from file")
	var dataPath = flag.String("load", "", "Load data to DB from file")
	flag.Parse()

	db, err := badger.Open(badger.DefaultOptions("./badger.db").WithCompression(options.ZSTD))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vdb = &VehicleDB{db}

	if *dataPath != "" {
		vdb.LoadData(*dataPath)
		return
	}

	// Port from ENV
	p := os.Getenv("PORT")
	if p == "" {
		p = "1337"
	}

	server(fmt.Sprintf("%s:%s", *host, p))
}
