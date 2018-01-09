package nflag

import (
	parser "gopkg.in/alecthomas/kingpin.v2"
)

func HandleNamespace(cmdHandler *parser.CmdClause) *string {
	return cmdHandler.Flag(
		"namespace",
		"optional namespace, e.g. `twilio`",
	).Short('n').String()
}

func HandleTags(cmdHandler *parser.CmdClause) *string {
	return cmdHandler.Flag(
		"tags",
		"comma separated list of tags for this note",
	).Short('t').String()
}
