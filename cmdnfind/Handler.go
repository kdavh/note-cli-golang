package cmdnfind

import (
	"fmt"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nflag"
	"github.com/kdavh/note-cli-golang/nparse"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/exec"
	"path/filepath"
	//"regexp"
	"github.com/kdavh/note-cli-golang/nflow"
	"strconv"
	"strings"
)

type Handler struct {
	handler   *parser.CmdClause
	tags      *string
	namespace *string
}

func (c *Handler) FullCommand() string {
	return c.handler.FullCommand()
}

func (c *Handler) Run(config nconfig.Config, ctx nctx.Context) bool {
	ctx.Logger.Debugf("SEARCH TAGS %v\n", nparse.CommaListtoa(*c.tags))
	notesPath := filepath.Join(os.Getenv("DOTFILES"), "notes")
	searchDepth := "0"

	var findGlobs []string
	if *c.namespace == "*" {
		// all namespaces
		findGlobs = []string{
			filepath.Join(notesPath),
		}
		searchDepth = "1"
	} else if *c.namespace == "" {
		// only root namespace
		findGlobs = []string{
			filepath.Join(notesPath),
		}
	} else {
		// namespace is specified
		for _, ns := range nparse.CommaListtoa(*c.namespace) {
			dir := filepath.Join(notesPath, ns)
			info, err := os.Stat(dir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else if !info.Mode().IsDir() {
				ctx.Logger.Errorf("\"%s\" is not a valid namespace (not a directory)\n", ns)
				os.Exit(1)
			}

			findGlobs = append(findGlobs, dir)
		}
	}
	ctx.Logger.Debugf("SEARCH DIRS: %v\n", findGlobs)

	var tagsLookaheads []string

	for _, tag := range nparse.CommaListtoa(*c.tags) {
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
		} else if len(files) == 0 {
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

func NewHandler(app *parser.Application) Handler {
	findNote := app.Command("find", "Find note.")

	findNoteTags := nflag.HandleTags(findNote)
	findNoteNamespace := nflag.HandleNamespace(findNote)

	return Handler{
		handler:   findNote,
		tags:      findNoteTags,
		namespace: findNoteNamespace,
	}
}
