package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/daflad/strava-reader/models"
	"github.com/daflad/strava-reader/models/app"
)

// runtime flags
var verbose bool
var version bool
var fileName string

func init() {
	// register flags
	flag.StringVar(&fileName, "p", "", "path to the Strava JSON file to parse")
	flag.BoolVar(&verbose, "v", false, "verbose logging flag")
	flag.BoolVar(&version, "V", false, "display app version")
}

func main() {
	// grab flag values
	flag.Parse()
	app.Init(verbose, version)
	if fileName == "" {
		// can't do much without a file to read
		flag.Usage()
		os.Exit(0)
	}
	// parse model from file
	stravaModel := models.ParseJSON(fileName)
	// build overview stats
	rideStats := models.RideStatsFromStrava(stravaModel)
	fmt.Println(rideStats)
}
