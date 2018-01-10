package nparse

import (
	"strings"
)

func CommaSplit(s string) []string {
	return strings.Split(s, ",")
}
