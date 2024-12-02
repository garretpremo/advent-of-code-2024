package util

import (
	"strings"
)

func ReadLines(bytes []byte) []string {
	return strings.Split(string(bytes), "\r\n")
}
