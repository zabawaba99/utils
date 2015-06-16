package stringy

import (
	"strings"
)

type String string

func (s String) Chomp(cutSet string) String {
	return String(strings.TrimRight(string(s), cutSet))
}

func (s String) String() string {
	return string(s)
}
