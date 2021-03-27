package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/daflad/strava-reader/models"
	"github.com/daflad/strava-reader/models/app"
	"github.com/wcharczuk/go-chart/v2"
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
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "The XAxis",
		},
		YAxis: chart.YAxis{
			Name: "The YAxis",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: rideStats.Distances,
				YValues: rideStats.Elevation,
			},
		},
	}

	f, err := os.Create("output.png")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	graph.Render(chart.PNG, f)
}
