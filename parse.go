package main

import (
	"strings"
)

func parseCommaList(s string) []string {
	return strings.Split(s, ",")
}
