package utils

import (
	"log"
	"os"
)

// Logger is the object that is used to log
type Logger interface {
	Printf(format string, args ...interface{})
}

// Log is current Logger
var Log Logger

func init() {
	Log = log.New(os.Stdout, "utils: ", log.Ltime)
}
