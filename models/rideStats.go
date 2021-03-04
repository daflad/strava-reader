package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// RideStats compiled for a given strave json file
type RideStats struct {
	RideName string
	Activity string
	RideDate string
	RideTime time.Duration
	Duration time.Duration
	Distance float64
	AvgSpeed float64
	AvgHrtRt float64
	FlatRide float64
	HillRide float64
	DownHill float64
	ElevGain float64
	ElevLoss float64
}

// RideStatsFromStrava build from json
func RideStatsFromStrava(s Strava) (re RideStats) {
	re.RideName = s.Name
	re.Activity = strings.Title(s.Metadata.ActivityType)
	re.RideDate = s.Metadata.StartDate.Format("Monday, January 2, 2006 @ 15:04:05")
	var err error
	re.RideTime, err = time.ParseDuration(strconv.Itoa(s.Metadata.ElapsedTime) + "s")
	if err != nil {
		log.Println(err)
	}
	re.Duration, err = time.ParseDuration(strconv.Itoa(s.Metadata.TimerTime) + "s")
	if err != nil {
		log.Println(err)
	}
	re.Distance = km2miles(s.Metadata.OverrideDistance)
	re.AvgSpeed = re.Distance / float64(re.Duration.Seconds()/60.0/60.0)
	re.AvgHrtRt = 0.0
	for _, hr := range s.Data[1].Values {
		re.AvgHrtRt += hr[1].(float64)
	}
	re.AvgHrtRt /= float64(len(s.Data[1].Values))
	re.ElevGain = 0.0
	re.ElevLoss = 0.0
	re.HillRide = 0.0
	re.DownHill = 0.0
	re.FlatRide = 0.0
	lastEl := 0.0
	lastDist := 0.0
	for i, el := range s.Data[0].Values {
		ele := el[2].(float64)
		dist := el[8].(float64)
		if i > 0 {
			if ele > lastEl {
				re.ElevGain += ele - lastEl
				re.HillRide += dist - lastDist
			} else if ele < lastEl {
				re.ElevLoss += ele - lastEl
				re.DownHill += dist - lastDist
			} else {
				re.FlatRide += dist - lastDist
			}
		}
		lastEl = ele
		lastDist = dist
	}
	return
}

func (re RideStats) String() string {
	formatString := `
RideName: %s
Activity: %s
RideDate: %s
RideTime: %s
Duration: %s
Distance: %.02f miles
AvgSpeed: %.02fmph
AvgHrtRt: %.02fbpm
FlatRide: %0.2f miles
HillRide: %0.2f miles
DownHill: %0.2f miles
ElevGain: %.02f meters
ElevLoss: %.02f meters
`
	return fmt.Sprintf(formatString, re.RideName, re.Activity, re.RideDate,
		re.Duration, re.RideTime, re.Distance, re.AvgSpeed, re.AvgHrtRt,
		km2miles(re.FlatRide), km2miles(re.HillRide), km2miles(re.DownHill),
		re.ElevGain, re.ElevLoss)
}

func km2miles(in interface{}) float64 {
	return in.(float64) * 0.000621
}

func m2miles(in interface{}) float64 {
	return in.(float64) * 0.00000621
}
