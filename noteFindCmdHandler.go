package main

import (
	"fmt"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/exec"
	"path/filepath"
	//"regexp"
	"strings"
)

type noteFindCmdHandler struct {
	handler   *parser.CmdClause
	tags      *string
	namespace *string
}

func (c *noteFindCmdHandler) FullCommand() string {
	return c.handler.FullCommand()
}

func (c *noteFindCmdHandler) Run(config AppConfig, ctx AppContext) bool {
	ctx.Logger.Debugf("%v\n", parseCommaList(*c.tags))
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
		for _, ns := range parseCommaList(*c.namespace) {
			dir := filepath.Join(notesPath, ns)
			info, err := os.Stat(dir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else if !info.Mode().IsDir() {
				fmt.Printf("\"%s\" is not a valid namespace (not a directory)\n", ns)
				os.Exit(1)
			}

			findGlobs = append(findGlobs, dir)
		}
	}
	fmt.Printf("dir globs: %v\n", findGlobs)

	var tagsLookaheads []string

	for _, tag := range parseCommaList(*c.tags) {
		tagsLookaheads = append(tagsLookaheads, fmt.Sprintf("(?=(.*,)?\\s*%s\\s*(,|$))", tag))
	}

	searchCmd := config.SearchApp + " \"" + TAGLINE + strings.Join(tagsLookaheads, "|") + "\" --files-with-matches --depth=" + searchDepth + " " + strings.Join(findGlobs, " ")
	ctx.Logger.Debugf("COMMAND: %s\n", searchCmd)

	if output, cmdErr := exec.Command("zsh", "-c", searchCmd).Output(); cmdErr != nil {
		if cmdErr.Error() == "exit status 1" {
			ctx.Logger.Error("No relevant files found")
		} else {
			ctx.Logger.Errorf("COMMAND FAILED: %s\nerror: %s", searchCmd, cmdErr)
		}

		os.Exit(1)
	} else {
		// TODO; prompt for file name most relevant, display
		fmt.Printf(string(output))
	}

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
