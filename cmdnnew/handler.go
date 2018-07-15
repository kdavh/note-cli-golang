package cmdnnew

import (
	"regexp"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/neditor"
	"github.com/kdavh/note-cli-golang/nflag"
	"github.com/kdavh/note-cli-golang/nflow"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type Handler struct {
	handler   *parser.CmdClause
	tags      *[]string
	fileName  *string
	namespace *string
	rp        nconfig.ReporterInterface
	ed        *neditor.Editor
	osCtrl    *nconfig.OsCtrl
}

func (hndl *Handler) CanHandle(commands []string) bool {
	return len(commands) > 0 && hndl.handler.FullCommand() == commands[0]
}

func (hndl *Handler) Run() bool {
	if match, _ := regexp.MatchString("\\.md$", *hndl.fileName); !match {
		hndl.rp.Errorf("%s must end with `.md`, exiting\n", *hndl.fileName)
		hndl.osCtrl.Exit(1)
	}

	hndl.rp.Debugf("%v\n", *hndl.tags)
	if err := hndl.ed.NewFile(*hndl.namespace, *hndl.fileName, *hndl.tags); err != nil {
		nflow.ErrExit(err, hndl.rp)
	}

	if err := hndl.ed.Open(*hndl.namespace, *hndl.fileName); err != nil {
		hndl.rp.Errorf(err.Error())
		hndl.osCtrl.Exit(1)
	} else {
		hndl.osCtrl.Exit(0)
	}

	return true
}

func NewHandler(app *parser.Application, rp nconfig.ReporterInterface, ed *neditor.Editor, osCtrl *nconfig.OsCtrl) *Handler {
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
		rp:        rp,
		ed:        ed,
		osCtrl:    osCtrl,
	}
}
