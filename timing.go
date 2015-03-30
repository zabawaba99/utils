package utils

import (
	"fmt"
	"time"
)

// Time gets the difference between the current time
// and the given start time and then logs the difference
//
//		defer Time(time.Now(), "some-code")
func Time(start time.Time, name string) {
	totalTime := time.Now().Sub(start)
	Writer.Write([]byte(fmt.Sprintf("%s took %s", name, totalTime)))
}
