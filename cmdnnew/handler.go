package cmdnnew

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nflag"
	"github.com/kdavh/note-cli-golang/nflow"
	"github.com/kdavh/note-cli-golang/nparse"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type Handler struct {
	handler   *parser.CmdClause
	tags      *string
	fileName  *string
	namespace *string
	config    *nconfig.Config
}

func (c *Handler) CanHandle(commands []string) bool {
	return len(commands) > 0 && c.handler.FullCommand() == commands[0]
}

func (c *Handler) Run() bool {
	cfg := c.config
	fs := cfg.Fs
	os := cfg.OsCtrl

	if match, _ := regexp.MatchString("\\.md$", *c.fileName); !match {
		cfg.Reporter.Errorf("%s must end with `.md`, exiting\n", *c.fileName)
		os.Exit(1)
	}

	cfg.Reporter.Debugf("%v\n", nparse.CommaSplit(*c.tags))
	newNoteDir := filepath.Join(cfg.NotesPath, *c.namespace)
	fs.MkdirAll(newNoteDir, 0755)
	filename := filepath.Join(newNoteDir, *c.fileName)
	if _, statErr := fs.Stat(filename); os.IsExist(statErr) {
		nflow.ErrExit(statErr, cfg.Reporter)
	} else {
		cfg.Reporter.Infof("CREATING FILE: %s\n", filename)

		file, newFileErr := fs.Create(filename)
		nflow.ErrExit(newFileErr, cfg.Reporter)

		data := cfg.Tagline + " " + strings.Join(nparse.CommaSplit(*c.tags), " ")
		fmt.Fprintf(file, data)
		file.Close()

		nflow.ShellOpen(filename, cfg)
	}

	return true
}

func NewHandler(app *parser.Application, config *nconfig.Config) *Handler {
	newNote := app.Command("new", "New note.")

	newNoteName := newNote.Arg(
		"name",
		"Name of note file, must end in `.md`.",
	).Required().String()

	newNoteTags := nflag.HandleTags(newNote)
	newNoteNamespace := nflag.HandleNamespace(newNote)

	return &Handler{
		handler:   newNote,
		tags:      newNoteTags,
		fileName:  newNoteName,
		namespace: newNoteNamespace,
		config:    config,
	}
}
