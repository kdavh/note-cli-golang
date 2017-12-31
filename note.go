package main

import (
	"fmt"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const TAGLINE = "###-tags-:"

func main() {
	var app = parser.New("note", "A command-line note keeping application with tags.")
	app.HelpFlag.Short('h')

	var (
		//debug    = app.Flag("debug", "Enable debug mode.").Bool()
		//serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

		newNote = app.Command("new", "New note.")

		newNoteName = newNote.Arg("name", "Name of note file, must end in `.md`.").Required().String()

		post      = app.Command("post", "Post a message to a channel.")
		postImage = post.Flag("image", "Image to post.").File()
		//postChannel = post.Arg("channel", "Channel to post to.").Required().String()
		postText = post.Arg("text", "Text to post.").Strings()
	)

	var (
		newNoteTagsFlagExpl = "comma separated list of tags for this note"
		newNoteTagsFlag     = newNote.Flag("tags", newNoteTagsFlagExpl).Short('t')
		newNoteTags         = newNoteTagsFlag.String()
	)

	var (
		newNoteNsFlagExpl = "optional namespace for the flag, e.g. `twilio`"
		newNoteNsFlag     = newNote.Flag("namespace", newNoteNsFlagExpl).Short('n')
		newNoteNs         = newNoteNsFlag.String()
	)

	// parser fills in values of flags and args here
	switch parser.MustParse(app.Parse(os.Args[1:])) {
	// new note
	case newNote.FullCommand():
		if match, _ := regexp.MatchString("\\.md$", *newNoteName); !match {
			fmt.Printf("%s must end with `.md`, exiting\n", *newNoteName)
			os.Exit(1)
		}

		fmt.Printf("%v\n", parseCommaList(*newNoteTags))
		newNoteDir := filepath.Join(os.Getenv("DOTFILES"), "notes", *newNoteNs)
		os.MkdirAll(newNoteDir, 0755)
		filename := filepath.Join(newNoteDir, *newNoteName)
		if _, statErr := os.Stat(filename); os.IsExist(statErr) {
			errExit(statErr)
		} else {
			fmt.Printf("creating %s\n", filename)

			file, newFileErr := os.Create(filename)
			errExit(newFileErr)
			defer file.Close()

			data := TAGLINE + " " + strings.Join(parseCommaList(*newNoteTags), ", ")
			fmt.Fprintf(file, data)
		}

	// Post message
	case post.FullCommand():
		if *postImage != nil {
		}
		text := strings.Join(*postText, " ")
		println("Post:", text)
	}
}
