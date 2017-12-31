package main

import (
	"fmt"
	parser "gopkg.in/alecthomas/kingpin.v2"
	//"os"
	//"path/filepath"
	//"regexp"
	//"strings"
)

type noteFindCmdHandler struct {
	handler   *parser.CmdClause
	tags      *string
	namespace *string
}

func (c *noteFindCmdHandler) FullCommand() string {
	return c.handler.FullCommand()
}

func (c *noteFindCmdHandler) Run() bool {
	fmt.Printf("%v\n", parseCommaList(*c.tags))

	if *c.namespace == "" {
	} else {
	}
	//findNoteDir := filepath.Join(os.Getenv("DOTFILES"), "notes", *c.Ns)
	//filename := filepath.Join(newNoteDir, *c.FileName)
	//if _, statErr := os.Stat(filename); os.IsExist(statErr) {
	//errExit(statErr)
	//} else {
	//fmt.Printf("creating %s\n", filename)

	//file, newFileErr := os.Create(filename)
	//errExit(newFileErr)
	//defer file.Close()

	//data := TAGLINE + " " + strings.Join(parseCommaList(*c.Tags), ", ")
	//fmt.Fprintf(file, data)
	//}

	return true
}

func createNoteFindCmdHandler(app *parser.Application) noteFindCmdHandler {
	findNote := app.Command("find", "Find note.")

	findNoteTags := handleTagsFlag(findNote)
	findNoteNamespace := handleNamespaceFlag(findNote)

	return noteFindCmdHandler{
		handler:   findNote,
		tags:      findNoteTags,
		namespace: findNoteNamespace,
	}
}
