package utils

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	testName := "cool test"
	Time(time.Now(), testName)

	assert.Len(t, writer.entries, 1)
	assert.Regexp(t, regexp.MustCompile("^"+testName+" took "), writer.entries[0])
}
