package main

import (
	//"fmt"
	"github.com/kdavh/note-cli-golang/cmdnfind"
	"github.com/kdavh/note-cli-golang/cmdnnew"
	"github.com/kdavh/note-cli-golang/cmdtag"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nlog"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

func main() {
	appContext := nctx.Context{
		Logger: nlog.New(nlog.ERROR),
	}

	appConfig := nconfig.Config{
		SearchApp: "ag",
		Editor:    "nvim",
		Tagline:   "###-tags-:",
	}

	app := parser.New("note", "A command-line note keeping application with tags.")
	app.HelpFlag.Short('h')

	verbose := app.Flag("verbose", "Enable debug mode.").Short('v').Bool()

	noteNewCmdHandler := cmdnnew.NewHandler(app)
	noteFindCmdHandler := cmdnfind.NewHandler(app)

	tagCmdHandler := cmdtag.NewHandler(app)
	var (
	//debug    = app.Flag("debug", "Enable debug mode.").Bool()
	//serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

	//post      = app.Command("post", "Post a message to a channel.")
	//postImage = post.Flag("image", "Image to post.").File()
	//postChannel = post.Arg("channel", "Channel to post to.").Required().String()
	//postText = post.Arg("text", "Text to post.").Strings()
	)

	// parser fills in values of flags and args here
	subcommands := strings.Split(parser.MustParse(app.Parse(os.Args[1:])), " ")
	if *verbose {
		appContext.Logger = nlog.New(nlog.DEBUG)
	}

	switch subcommands[0] {
	case noteNewCmdHandler.FullCommand():
		noteNewCmdHandler.Run(appConfig, appContext)
	case noteFindCmdHandler.FullCommand():
		noteFindCmdHandler.Run(appConfig, appContext)
	case tagCmdHandler.FullCommand():
		tagCmdHandler.Run(appConfig, appContext, subcommands)
	}
}
