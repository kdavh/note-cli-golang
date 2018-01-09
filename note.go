package main

import (
	"github.com/kdavh/note-cli-golang/nlog"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const TAGLINE = "###-tags-:"

type AppConfig struct {
	SearchApp string
	Editor    string
}

type AppContext struct {
	Logger *nlog.Logger
}

func main() {
	appContext := AppContext{
		Logger: nlog.New(nlog.DEBUG),
	}

	appConfig := AppConfig{
		SearchApp: "ag",
		Editor:    "nvim",
	}

	app := parser.New("note", "A command-line note keeping application with tags.")
	app.HelpFlag.Short('h')

	noteNewCmdHandler := createNoteNewCmdHandler(app)
	noteFindCmdHandler := createNoteFindCmdHandler(app)
	var (
	//debug    = app.Flag("debug", "Enable debug mode.").Bool()
	//serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

	//post      = app.Command("post", "Post a message to a channel.")
	//postImage = post.Flag("image", "Image to post.").File()
	//postChannel = post.Arg("channel", "Channel to post to.").Required().String()
	//postText = post.Arg("text", "Text to post.").Strings()
	)

	// parser fills in values of flags and args here
	switch parser.MustParse(app.Parse(os.Args[1:])) {
	// new note
	case noteNewCmdHandler.FullCommand():
		noteNewCmdHandler.Run(appConfig, appContext)
	case noteFindCmdHandler.FullCommand():
		noteFindCmdHandler.Run(appConfig, appContext)
		// Post message
		//case post.FullCommand():
		//if *postImage != nil {
		//}
		//text := strings.Join(*postText, " ")
		//println("Post:", text)
	}
}
