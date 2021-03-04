package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/daflad/strava-reader/models/app"
)

// Strava struct to accept json from strava.settings.ExportOriginal
type Strava struct {
	Name     string
	Metadata struct {
		HTTPUserAgent      string    `json:"http_user_agent"`
		AthleteID          int       `json:"athlete_id"`
		ExternalID         string    `json:"external_id"`
		DataType           string    `json:"data_type"`
		ActivityType       string    `json:"activity_type"`
		Private            bool      `json:"private"`
		StartDate          time.Time `json:"start_date"`
		TimerTime          int       `json:"timer_time"`
		ElapsedTime        int       `json:"elapsed_time"`
		MinAccuracy        int       `json:"min_accuracy"`
		DistanceFilter     int       `json:"distance_filter"`
		SampleRate         int       `json:"sample_rate"`
		TimeSeriesField    string    `json:"time_series_field"`
		WorkoutType        int       `json:"workout_type"`
		Commute            bool      `json:"commute"`
		ScreenOnDuration   int       `json:"screen_on_duration"`
		InitialBatteryLife float64   `json:"initial_battery_life"`
		FinalBatteryLife   float64   `json:"final_battery_life"`
		AutopauseEnabled   bool      `json:"autopause_enabled"`
		PreventAutolock    bool      `json:"prevent_autolock"`
		LiveActivityID     int       `json:"live_activity_id"`
		OverrideDistance   float64   `json:"override_distance"`
	} `json:"metadata"`
	Data []struct {
		Values [][]interface{} `json:"values"`
		Fields []string        `json:"fields"`
	} `json:"data"`
}

// ParseJSON file from path and return the Strava object
func ParseJSON(path string) (ride Strava) {
	// open file and read all lines into struct
	jsonFile, err := os.Open(path)
	app.CheckForError(err, "Not a good file path", "ParseJSON()")
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	app.CheckForError(err, "Error reading lines from file", "ParseJSON()")
	err = json.Unmarshal(data, &ride)
	app.CheckForError(err, "JSON formatting error", "ParseJSON()")
	// grab the name of the ride from the path /path/to/my/FILENAME.json
	quickName := path[strings.LastIndex(path, "/")+1 : strings.LastIndex(path, ".")]
	ride.Name = strings.Replace(quickName, "_", " ", -1)
	return
}
