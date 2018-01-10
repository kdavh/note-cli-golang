package cmdnfind

import (
	"fmt"
	"github.com/kdavh/note-cli-golang/cmdhelp"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nflag"
	"github.com/kdavh/note-cli-golang/nparse"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/exec"
	//"regexp"
	"github.com/kdavh/note-cli-golang/nflow"
	"strconv"
	"strings"
)

type Handler struct {
	handler   *parser.CmdClause
	tags      *string
	namespace *string
	config    *nconfig.Config
	ctx       *nctx.Context
}

func (c *Handler) CanHandle(commands []string) bool {
	return len(commands) > 0 && c.handler.FullCommand() == commands[0]
}

func (c *Handler) Run() bool {
	ctx := c.ctx
	config := c.config

	ctx.Logger.Debugf("SEARCH TAGS %v\n", nparse.CommaSplit(*c.tags))

	findGlobs, searchDepth := cmdhelp.FileGlobs(*c.namespace, config, ctx)

	var tagsLookaheads []string

	for _, tag := range nparse.CommaSplit(*c.tags) {
		tagsLookaheads = append(tagsLookaheads, fmt.Sprintf("(?=\\s+%s(\\s+|$))", tag))
	}

	searchCmd := config.SearchApp + " \"" + config.Tagline + strings.Join(tagsLookaheads, "|") + "\" --files-with-matches --depth=" + searchDepth + " " + strings.Join(findGlobs, " ")
	ctx.Logger.Debugf("SEARCH COMMAND: %s\n", searchCmd)

	if output, cmdErr := exec.Command("zsh", "-c", searchCmd).Output(); cmdErr != nil {
		if cmdErr.Error() == "exit status 1" {
			ctx.Logger.Error("No relevant files found")
		} else {
			ctx.Logger.Errorf("COMMAND FAILED: %s\nerror: %s", searchCmd, cmdErr)
		}

		os.Exit(1)
	} else {
		var chosenFile string
		// TODO; prompt for file name most relevant, display
		fmt.Printf(string(output))
		files := strings.Split(strings.TrimSpace(string(output)), "\n")

		if len(files) == 0 {
			ctx.Logger.Error("Should never reach here... blurg")
			os.Exit(1)
		} else if len(files) == 1 {
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

		nflow.ShellOpen(config.Editor, chosenFile, ctx.Logger)
	}

	return true
}

func NewHandler(app *parser.Application, config *nconfig.Config, ctx *nctx.Context) Handler {
	findNote := app.Command("find", "Find note.")

	findNoteTags := nflag.HandleTags(findNote)
	findNoteNamespace := nflag.HandleNamespace(findNote)

	return Handler{
		handler:   findNote,
		tags:      findNoteTags,
		namespace: findNoteNamespace,
		config:    config,
		ctx:       ctx,
	}
}
