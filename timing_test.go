package utils

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	testName := "cool test"
	Time(time.Now(), testName)

	tl, ok := Log.(*testLogger)
	require.True(t, ok)

	assert.Len(t, tl.entries, 1)
	assert.Regexp(t, regexp.MustCompile("^"+testName+" took "), tl.entries[0])
}
