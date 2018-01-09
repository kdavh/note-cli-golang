package main

import (
	"fmt"
	"github.com/kdavh/note-cli-golang/nflow"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type noteNewCmdHandler struct {
	handler   *parser.CmdClause
	tags      *string
	fileName  *string
	namespace *string
}

func (c *noteNewCmdHandler) FullCommand() string {
	return c.handler.FullCommand()
}

func (c *noteNewCmdHandler) Run(config AppConfig, ctx AppContext) bool {
	if match, _ := regexp.MatchString("\\.md$", *c.fileName); !match {
		fmt.Printf("%s must end with `.md`, exiting\n", *c.fileName)
		os.Exit(1)
	}

	fmt.Printf("%v\n", parseCommaList(*c.tags))
	newNoteDir := filepath.Join(os.Getenv("DOTFILES"), "notes", *c.namespace)
	os.MkdirAll(newNoteDir, 0755)
	filename := filepath.Join(newNoteDir, *c.fileName)
	if _, statErr := os.Stat(filename); os.IsExist(statErr) {
		nflow.ErrExit(statErr, ctx.Logger)
	} else {
		fmt.Printf("creating %s\n", filename)

		file, newFileErr := os.Create(filename)
		nflow.ErrExit(newFileErr, ctx.Logger)

		data := TAGLINE + " " + strings.Join(parseCommaList(*c.tags), " ")
		fmt.Fprintf(file, data)
		file.Close()

		nflow.ShellOpen(config.Editor, filename, ctx.Logger)
	}

	return true
}

func createNoteNewCmdHandler(app *parser.Application) noteNewCmdHandler {
	newNote := app.Command("new", "New note.")

	newNoteName := newNote.Arg(
		"name",
		"Name of note file, must end in `.md`.",
	).Required().String()

	newNoteTags := handleTagsFlag(newNote)

	newNoteNamespace := handleNamespaceFlag(newNote)

	return noteNewCmdHandler{
		handler:   newNote,
		tags:      newNoteTags,
		fileName:  newNoteName,
		namespace: newNoteNamespace,
	}
}
