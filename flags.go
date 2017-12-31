package main

import (
	parser "gopkg.in/alecthomas/kingpin.v2"
)

func handleNamespaceFlag(cmdHandler *parser.CmdClause) *string {
	return cmdHandler.Flag(
		"namespace",
		"optional namespace, e.g. `twilio`",
	).Short('n').String()
}

func handleTagsFlag(cmdHandler *parser.CmdClause) *string {
	return cmdHandler.Flag(
		"tags",
		"comma separated list of tags for this note",
	).Short('t').String()
}
