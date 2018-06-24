package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"github.com/kdavh/note-cli-golang/cmdnfind"
	"github.com/kdavh/note-cli-golang/cmdnnew"
	"github.com/kdavh/note-cli-golang/cmdtag"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nreport"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	devDir := filepath.Join(os.Getenv("HOME"), "dev")
	appConfig := &nconfig.Config{
		SearchApp:    "ag",
		Editor:       "nvim",
		EditorConfig: filepath.Join(devDir, "note-app-vim", "vim-note-config.vimrc"),
		Tagline:      "###-tags-:",
		NotesPath:    filepath.Join(devDir, "note-app-notes", "notes"),
		Fs:           afero.NewOsFs(),
		OsCtrl:       nconfig.NewOsCtrl(),
		Reporter:     nreport.New(nreport.INFO),
	}

	// configure the parser flags and subcommands
	app := parser.New("note", "A command-line note keeping application with tags.")
	app.HelpFlag.Short('h')

	verbose := app.Flag("verbose", "Enable debug mode.").Short('v').Bool()

	noteNewCmdHandler := cmdnnew.NewHandler(app, appConfig)
	noteFindCmdHandler := cmdnfind.NewHandler(app, appConfig)
	tagCmdHandler := cmdtag.NewHandler(app, appConfig)

	// parser fills in values of flags, commands and returns subcommands here
	commands := strings.Split(parser.MustParse(app.Parse(os.Args[1:])), " ")

	// starts responding to command
	if *verbose {
		appConfig.Reporter = nreport.New(nreport.DEBUG)
	}
	appConfig.Reporter.Debugf("%v\n\n", commands)

	if noteNewCmdHandler.CanHandle(commands) {
		noteNewCmdHandler.Run()
	} else if noteFindCmdHandler.CanHandle(commands) {
		noteFindCmdHandler.Run()
	} else if tagCmdHandler.CanHandle(commands) {
		tagCmdHandler.Run(commands)
	}
}
