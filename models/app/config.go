package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Config struct to hold app settings
type Config struct {
	AppName    string       `json:"appName"`
	AppVersion string       `json:"appVersion"`
	DataDir    string       `json:"dataDir"`
	Database   DbConnection `json:"database"`
}

func (c *Config) version() {
	Info("%v version %v", c.AppName, c.AppVersion)
	os.Exit(0)
}

//DbConnection details
type DbConnection struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

//ConnectionString for the db
func (d *DbConnection) ConnectionString(display bool) string {
	pass := d.Pass
	if display {
		pass = "****"
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s",
		d.User, pass, "tcp", d.Host, d.Port, d.Name)
}

func loadConfig(printVersion bool) (c *Config) {
	dat, err := ioutil.ReadFile("config.json")
	if err != nil {
		CheckForError(err, "Please provide a valid config file", "LoadConfig()")
	}
	if err := json.Unmarshal(dat, &c); err != nil {
		CheckForError(err, "JSON UnMarshal", "LoadConfig()")
	}
	InfoVebose("%v Data Parsing & Analysis", c.AppName)
	if printVersion {
		c.version()
	}
	return c
}
