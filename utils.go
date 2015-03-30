package utils

import (
	"io"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "utils: ", log.Ldate)
}

// SetWriter configures the logging that occurs in this function
// the parameters that are passed to the functions are passed directly
// to the log package
//
//		logger = log.New(w, prefix, flag)
func SetWriter(w io.Writer, prefix string, flag int) {
	logger = log.New(w, prefix, flag)
}
