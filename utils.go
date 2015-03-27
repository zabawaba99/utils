package utils

import (
	"log"
	"os"
)

type Logger interface {
	Printf(format string, args ...interface{})
}

var Log Logger

func init() {
	Log = log.New(os.Stdout, "utils: ", log.Ltime)
}
