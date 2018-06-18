package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"github.com/kdavh/note-cli-golang/cmdnfind"
	"github.com/kdavh/note-cli-golang/cmdnnew"
	"github.com/kdavh/note-cli-golang/cmdtag"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nlog"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	appContext := &nctx.Context{
		Logger: nlog.New(nlog.ERROR),
	}

	devDir := filepath.Join(os.Getenv("HOME"), "dev")
	appConfig := &nconfig.Config{
		SearchApp:    "ag",
		Editor:       "nvim",
		EditorConfig: filepath.Join(devDir, "note-app-vim", "vim-note-config.vimrc"),
		Tagline:      "###-tags-:",
		NotesPath:    filepath.Join(devDir, "note-app-notes", "notes"),
		Fs:           afero.NewOsFs(),
		OsCtrl: nconfig.OsCtrl{
			os.Exit,
		},
	}

	app := parser.New("note", "A command-line note keeping application with tags.")
	app.HelpFlag.Short('h')

	verbose := app.Flag("verbose", "Enable debug mode.").Short('v').Bool()

	noteNewCmdHandler := cmdnnew.NewHandler(app, appConfig, appContext)
	noteFindCmdHandler := cmdnfind.NewHandler(app, appConfig, appContext)
	tagCmdHandler := cmdtag.NewHandler(app, appConfig, appContext)

	// parser fills in values of flags and args here
	commands := strings.Split(parser.MustParse(app.Parse(os.Args[1:])), " ")
	fmt.Printf("%v\n\n", commands)
	if *verbose {
		appContext.Logger = nlog.New(nlog.DEBUG)
	}

	if noteNewCmdHandler.CanHandle(commands) {
		noteNewCmdHandler.Run()
	} else if noteFindCmdHandler.CanHandle(commands) {
		noteFindCmdHandler.Run()
	} else if tagCmdHandler.CanHandle(commands) {
		tagCmdHandler.Run(commands)
	}
}
