package cmdnnew

import (
	"fmt"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nflag"
	"github.com/kdavh/note-cli-golang/nflow"
	"github.com/kdavh/note-cli-golang/nparse"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Handler struct {
	handler   *parser.CmdClause
	tags      *string
	fileName  *string
	namespace *string
}

func (c *Handler) FullCommand() string {
	return c.handler.FullCommand()
}

func (c *Handler) Run(config nconfig.Config, ctx nctx.Context) bool {
	if match, _ := regexp.MatchString("\\.md$", *c.fileName); !match {
		fmt.Printf("%s must end with `.md`, exiting\n", *c.fileName)
		os.Exit(1)
	}

	fmt.Printf("%v\n", nparse.CommaListtoa(*c.tags))
	newNoteDir := filepath.Join(os.Getenv("DOTFILES"), "notes", *c.namespace)
	os.MkdirAll(newNoteDir, 0755)
	filename := filepath.Join(newNoteDir, *c.fileName)
	if _, statErr := os.Stat(filename); os.IsExist(statErr) {
		nflow.ErrExit(statErr, ctx.Logger)
	} else {
		fmt.Printf("CREATING FILE: %s\n", filename)

		file, newFileErr := os.Create(filename)
		nflow.ErrExit(newFileErr, ctx.Logger)

		data := config.Tagline + " " + strings.Join(nparse.CommaListtoa(*c.tags), " ")
		fmt.Fprintf(file, data)
		file.Close()

		nflow.ShellOpen(config.Editor, filename, ctx.Logger)
	}

	return true
}

func NewHandler(app *parser.Application) Handler {
	newNote := app.Command("new", "New note.")

	newNoteName := newNote.Arg(
		"name",
		"Name of note file, must end in `.md`.",
	).Required().String()

	newNoteTags := nflag.HandleTags(newNote)
	newNoteNamespace := nflag.HandleNamespace(newNote)

	return Handler{
		handler:   newNote,
		tags:      newNoteTags,
		fileName:  newNoteName,
		namespace: newNoteNamespace,
	}
}
