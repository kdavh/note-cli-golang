package main

import (
	"os"

	"github.com/spf13/afero"

	"github.com/kdavh/note-cli-golang/cmdnamespace"
	"github.com/kdavh/note-cli-golang/cmdnfind"
	"github.com/kdavh/note-cli-golang/cmdnnew"
	"github.com/kdavh/note-cli-golang/cmdtag"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/neditor"
	"github.com/kdavh/note-cli-golang/nreport"
	"github.com/kdavh/note-cli-golang/nsearch"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	fs := afero.NewOsFs()
	ed := neditor.NewEditorVim(fs)
	osCtrl := nconfig.NewOsCtrl()
	rp := nreport.New(nreport.ERROR)
	se := nsearch.NewSearcherAg(fs)

	// configure the parser flags and subcommands
	app := parser.New("note", "A command-line note keeping application with tags.")
	app.HelpFlag.Short('h')

	verbose := app.Flag("verbose", "Enable debug mode.").Short('v').Bool()

	noteNewCmdHandler := cmdnnew.NewHandler(app, rp, ed, osCtrl)
	noteFindCmdHandler := cmdnfind.NewHandler(app, se, ed, osCtrl, rp)
	tagCmdHandler := cmdtag.NewHandler(app, se, osCtrl, rp)
	namespaceCmdHandler := cmdnamespace.NewHandler(app, se, osCtrl, rp)

	// parser fills in values of flags, commands and returns subcommands here
	commands := parser.MustParse(app.Parse(os.Args[1:]))

	// starts responding to command
	if *verbose {
		rp.Level = nreport.DEBUG
	}
	rp.Debugf("%v\n\n", commands)

	if noteNewCmdHandler.CanHandle(commands) {
		noteNewCmdHandler.Run()
	} else if noteFindCmdHandler.CanHandle(commands) {
		noteFindCmdHandler.Run()
	} else if tagCmdHandler.CanHandle(commands) {
		tagCmdHandler.Run(commands)
	} else if namespaceCmdHandler.CanHandle(commands) {
		namespaceCmdHandler.Run(commands)
	}
}
