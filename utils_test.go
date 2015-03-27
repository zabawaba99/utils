package utils

import (
	"fmt"
	"os"
	"testing"
)

// check compilation
var _ Logger = &testLogger{}

type testLogger struct {
	entries []string
}

func (tl *testLogger) Printf(format string, args ...interface{}) {
	tl.entries = append(tl.entries, fmt.Sprintf(format, args...))
}

func TestMain(m *testing.M) {
	Log = &testLogger{}
	os.Exit(m.Run())
}
