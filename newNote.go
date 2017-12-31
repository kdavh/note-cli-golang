package main

import (
	"fmt"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type noteNewCmdHandler struct {
	Handler  *parser.CmdClause
	Tags     *string
	FileName *string
	Ns       *string
}

func (c *noteNewCmdHandler) FullCommand() string {
	return c.Handler.FullCommand()
}

func (c *noteNewCmdHandler) Run() bool {
	if match, _ := regexp.MatchString("\\.md$", *c.FileName); !match {
		fmt.Printf("%s must end with `.md`, exiting\n", *c.FileName)
		os.Exit(1)
	}

	fmt.Printf("%v\n", parseCommaList(*c.Tags))
	newNoteDir := filepath.Join(os.Getenv("DOTFILES"), "notes", *c.Ns)
	os.MkdirAll(newNoteDir, 0755)
	filename := filepath.Join(newNoteDir, *c.FileName)
	if _, statErr := os.Stat(filename); os.IsExist(statErr) {
		errExit(statErr)
	} else {
		fmt.Printf("creating %s\n", filename)

		file, newFileErr := os.Create(filename)
		errExit(newFileErr)
		defer file.Close()

		data := TAGLINE + " " + strings.Join(parseCommaList(*c.Tags), ", ")
		fmt.Fprintf(file, data)
	}

	return true
}

func createNoteNewCmdHandler(app *parser.Application) noteNewCmdHandler {
	newNote := app.Command("new", "New note.")

	newNoteName := newNote.Arg("name", "Name of note file, must end in `.md`.").Required().String()

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

	return noteNewCmdHandler{
		Handler:  newNote,
		Tags:     newNoteTags,
		FileName: newNoteName,
		Ns:       newNoteNs,
	}
}
