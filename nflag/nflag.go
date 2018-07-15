package nflag

import (
	parser "gopkg.in/alecthomas/kingpin.v2"

	"github.com/kdavh/note-cli-golang/nparse"
)

func HandleNamespace(cmdHandler *parser.CmdClause) *string {
	return cmdHandler.Flag(
		"namespace",
		"optional namespace, e.g. `twilio`",
	).Short('n').String()
}

func HandleTags(cmdHandler *parser.CmdClause) *[]string {
	var cl *commaList = &commaList{}

	cmdHandler.Flag(
		"tags",
		"comma separated list of tags for this note",
	).Short('t').SetValue(cl)

	return (*[]string)(cl)
}

type commaList []string

func (cl *commaList) Set(value string) error {
	var l []string = append(*cl, nparse.CommaSplit(value)...)
	*cl = commaList(l)
	return nil
}

func (tl *commaList) String() string {
	return ""
}
