package app

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

//Error logger for app
type Error struct {
	ID      int64 `db:"id"`
	Error   int64 `db:"error"`
	Message int64 `db:"message"`
}

func (e *Error) String() string {
	return fmt.Sprintf("%v, %v", e.Error, e.Message)
}

//Verbose if logging should be verbose
var Verbose bool

//Init the Verbose bool
func Init(verbose, printVersion bool) {
	Verbose = verbose
	loadConfig(printVersion)
}

//CheckForError check the given bool and print / storte the result
func CheckForError(err error, message, funcName string) {
	if err != nil {
		if !strings.Contains(err.Error(), "sql: no rows in result set") {
			color.Red("ERROR: %v", color.YellowString(err.Error()))
			color.Red("ERROR: %v", color.YellowString("%v, %v", message, funcName))
		}
	} else {
		if Verbose {
			color.Green("SUCCESS: %v", color.YellowString(message))
		}
	}
}

//Info display an INFO message
func Info(format string, message ...interface{}) {
	if Verbose {
		color.Cyan("INFO: %v", color.YellowString(format, message...))
	} else {
		color.Cyan(format, message...)
	}
}

//InfoVebose display an INFO message
func InfoVebose(format string, message interface{}) {
	if Verbose {
		color.Cyan("INFO: %v", color.YellowString(format, message))
	}
}
