package utils

import (
	"io"
	"os"
)

// Writer is the default writer for any logging that occurs
// defaults to os.Stdout
var Writer io.Writer

func init() {
	Writer = os.Stdout
}
