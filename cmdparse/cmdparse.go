package cmdparse

import (
	"fmt"
	"path/filepath"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nparse"
)

// accept 0 or more namespaces, returning applicable dirs to be searched using search util
func FileGlobs(namespace string, config *nconfig.Config, ctx *nctx.Context) ([]string, string) {
	notesPath := config.NotesPath
	fs := config.Fs
	os := config.OsCtrl
	searchDepth := "0"
	var findGlobs []string

	if namespace == "*" {
		// all namespaces
		findGlobs = []string{
			notesPath,
		}
		searchDepth = "1"
	} else if namespace == "" {
		// only root namespace
		findGlobs = []string{
			notesPath,
		}
	} else {
		// namespace is specified
		for _, ns := range nparse.CommaSplit(namespace) {
			dir := filepath.Join(notesPath, ns)
			info, err := fs.Stat(dir)
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
	ctx.Logger.Debugf("SEARCH DEPTH: %v\n", searchDepth)

	return findGlobs, searchDepth
}
