package util

import (
	"strings"
)

func ReadLines(bytes []byte) []string {
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
