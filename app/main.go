package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/estraier/tkrzw-go"
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

	dbm := tkrzw.NewDBM()
	dbm.Open("db.tkh", true,
		tkrzw.ParseParams("truncate=true,num_buckets=10000"))

	vdb = &VehicleDB{dbm}

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
