package cmdparse

import (
	"path/filepath"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nparse"
)

// accept 0 or more namespaces, returning applicable dirs to be searched using search util
func FileGlobs(namespace string, config *nconfig.Config) ([]string, string) {
	notesPath := config.NotesPath
	fs := config.Fs
	os := config.OsCtrl
	reporter := config.Reporter
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
				reporter.Error(err.Error())
				os.Exit(1)
			} else if !info.Mode().IsDir() {
				reporter.Errorf("\"%s\" is not a valid namespace (not a directory)\n", ns)
				os.Exit(1)
			}

			findGlobs = append(findGlobs, dir)
		}
	}
	reporter.Debugf("SEARCH DIRS: %v\n", findGlobs)
	reporter.Debugf("SEARCH DEPTH: %v\n", searchDepth)

	return findGlobs, searchDepth
}
