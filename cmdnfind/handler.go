package cmdnfind

import (
	"fmt"
	"os"

	"strconv"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nflag"
	"github.com/kdavh/note-cli-golang/nparse"
	parser "gopkg.in/alecthomas/kingpin.v2"

	"github.com/kdavh/note-cli-golang/nflow"
)

type Handler struct {
	handler   *parser.CmdClause
	tags      *string
	namespace *string
	open      *bool
	config    *nconfig.Config
}

func (c *Handler) CanHandle(commands []string) bool {
	return len(commands) > 0 && c.handler.FullCommand() == commands[0]
}

func (c *Handler) Run() bool {
	config := c.config

	config.Reporter.Debugf("SEARCH TAGS %v\n", nparse.CommaSplit(*c.tags))

	files, cmdErr := config.Searcher.Notes(*c.namespace, *c.tags, "", config)
	if cmdErr != nil {
		if cmdErr.Error() == "exit status 1" {
			config.Reporter.Error("No relevant files found")
		} else {
			config.Reporter.Errorf("COMMAND FAILED: %s", cmdErr)
		}

		config.OsCtrl.Exit(1)
	} else {
		if *c.open {
			var chosenFile string

			if len(files) == 1 {
				chosenFile = files[0]
			} else {
				fmt.Println("Choose an option:")
				for i, file := range files {
					fmt.Printf("%s) %s\n", strconv.Itoa(i+1), file)
				}

				var input string
				fmt.Scanln(&input)

				chosenNumber, choiceErr := strconv.Atoi(input)

				if choiceErr != nil || chosenNumber > len(files) {
					fmt.Printf("\"%s\" is not a valid choice!\n", input)
					os.Exit(1)
				}

				chosenFile = files[chosenNumber-1]
			}

			nflow.ShellOpen(chosenFile, config)
		} else {
			fmt.Println("FOUND:")
			for _, file := range files {
				fmt.Printf("\t%s\n", file)
			}
		}
	}

	return true
}

func NewHandler(app *parser.Application, config *nconfig.Config) *Handler {
	findNote := app.Command("find", "Find note.")

	findNoteOpen := findNote.Flag("open", "Open files instead of just printing to stdout").Short('o').Bool()
	findNoteTags := nflag.HandleTags(findNote)
	findNoteNamespace := nflag.HandleNamespace(findNote)

	return &Handler{
		handler:   findNote,
		tags:      findNoteTags,
		namespace: findNoteNamespace,
		open:      findNoteOpen,
		config:    config,
	}
}
