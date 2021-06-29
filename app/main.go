package main

import (
	"flag"
	"log"

	badger "github.com/dgraph-io/badger/v3"
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
		return c.JSON(404, jsonError{"Not found"})
	}

	return c.JSONPretty(200, vehicle, "    ")
}

func server() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/lookup/:plate", plateLookup)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func main() {
	var dataPath = flag.String("load", "", "Load data to DB from file")
	flag.Parse()

	db, err := badger.Open(badger.DefaultOptions("./database"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vdb = &VehicleDB{db}

	if *dataPath != "" {
		vdb.LoadData(*dataPath)
		return
	}

	server()
}
