package nparse

import (
	"strings"
)

func CommaListtoa(s string) []string {
	return strings.Split(s, ",")
}
