package utils

import (
	"io"
	"os"
	"testing"
)

// check compilation
var _ io.Writer = &testWriter{}

type testWriter struct {
	entries []string
}

func (tl *testWriter) Write(p []byte) (int, error) {
	tl.entries = append(tl.entries, string(p))
	return 0, nil
}

var writer = &testWriter{}

func TestMain(m *testing.M) {
	SetWriter(writer, "", 0)
	os.Exit(m.Run())
}
