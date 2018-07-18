package cmdnfind

import (
	"strconv"
	"strings"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/neditor"
	"github.com/kdavh/note-cli-golang/nflag"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

const NO_NOTES_FOUND = "No relevant files found"

type Handler struct {
	handler   *parser.CmdClause
	tags      *[]string
	namespace *string
	open      *bool
	se        nconfig.SearcherInterface
	ed        *neditor.Editor
	osCtrl    *nconfig.OsCtrl
	rp        nconfig.ReporterInterface
}

func (hndl *Handler) CanHandle(commands string) bool {
	return strings.HasPrefix(commands, hndl.handler.FullCommand())
}

func (hndl *Handler) Run() bool {
	rp := hndl.rp

	rp.Debugf("SEARCH TAGS %v\n", *hndl.tags)

	files, cmdErr := hndl.se.Notes(*hndl.namespace, *hndl.tags, "", rp)
	if cmdErr != nil {
		if cmdErr.Error() == "exit status 1" {
			rp.Errorf(NO_NOTES_FOUND + "\n")
		} else {
			rp.Errorf("COMMAND FAILED: %s", cmdErr)
		}

		hndl.osCtrl.Exit(1)
	} else {
		if *hndl.open {
			var chosenFile string

			if len(files) == 1 {
				chosenFile = files[0]
			} else {
				rp.Reportf("Choose an option:")
				for i, file := range files {
					rp.Reportf("%s) %s\n", strconv.Itoa(i+1), file)
				}

				input := rp.Prompt()

				chosenNumber, choiceErr := strconv.Atoi(input)

				if choiceErr != nil || chosenNumber > len(files) {
					rp.Errorf("\"%s\" is not a valid choice!\n", input)
					hndl.osCtrl.Exit(1)
				}

				chosenFile = files[chosenNumber-1]
			}

			if err := hndl.ed.Open(*hndl.namespace, chosenFile); err != nil {
				hndl.osCtrl.Exit(1)
			} else {
				hndl.osCtrl.Exit(0)
			}
		} else {
			rp.Reportf("FOUND:\n")
			for _, file := range files {
				rp.Reportf("\t%s\n", file)
			}
		}
	}

	return true
}

func NewHandler(app *parser.Application, se nconfig.SearcherInterface, ed *neditor.Editor, osCtrl *nconfig.OsCtrl, rp nconfig.ReporterInterface) *Handler {
	findNote := app.Command("find", "Find note.")

	findNoteOpen := findNote.Flag("open", "Open files instead of just printing to stdout").Short('o').Bool()
	findNoteTags := nflag.HandleTags(findNote)
	findNoteNamespace := nflag.HandleNamespace(findNote)

	return &Handler{
		handler:   findNote,
		tags:      findNoteTags,
		namespace: findNoteNamespace,
		open:      findNoteOpen,
		ed:        ed,
		se:        se,
		osCtrl:    osCtrl,
		rp:        rp,
	}
}
