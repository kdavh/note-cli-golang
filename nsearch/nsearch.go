package nsearch

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nparse"
	"github.com/spf13/afero"
)

type searcher struct {
	prog          string
	tagLine       string
	notesDirPath  string
	tagLineFormat string
	fs            afero.Fs
}

func (se *searcher) Notes(namespace string, tagsQuery []string, textQuery string, rp nconfig.ReporterInterface) ([]string, error) {
	fileGlobs, searchDepth, err := se.fileGlobs(namespace)

	rp.Debugf("SEARCH DIRS: %v\n", fileGlobs)
	rp.Debugf("SEARCH DEPTH: %v\n", searchDepth)

	if err != nil {
		return []string{}, err
	}

	var tagsLookaheads []string
	for _, tag := range tagsQuery {
		tagsLookaheads = append(tagsLookaheads, fmt.Sprintf("(?=.*\\s+%s(\\s+|\\$))", tag))
	}

	cmd := exec.Command(se.prog, append([]string{
		se.tagLineFormat + strings.Join(tagsLookaheads, "|"),
		"--files-with-matches",
		"--depth=" + searchDepth,
	}, fileGlobs...)...)

	rp.Debugf("SEARCH COMMAND: %s\n", strings.Join(cmd.Args, " "))

	if output, cmdErr := cmd.Output(); cmdErr != nil {
		return []string{}, cmdErr
	} else {
		return strings.Split(strings.TrimSpace(string(output)), "\n"), nil
	}
}

func NewSearcherAg(fs afero.Fs) *searcher {
	return &searcher{
		prog:          "ag",
		notesDirPath:  nconfig.NotesDirPath(),
		tagLineFormat: nconfig.DefaultTaglineFormat(),
		fs:            fs,
	}
}

func NewSearcherMock(fs afero.Fs) *searcher {
	return &searcher{
		prog:          "ag",
		notesDirPath:  nconfig.NotesDirMockPath(),
		tagLineFormat: nconfig.DefaultTaglineFormat(),
		fs:            fs,
	}
}

func (se *searcher) fileGlobs(namespace string) ([]string, string, error) {
	notesPath := se.notesDirPath
	fs := se.fs

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
				return []string{}, "", err
				// reporter.Error(err.Error())
				// os.Exit(1)
			} else if !info.Mode().IsDir() {
				return []string{}, "", errors.New(fmt.Sprintf("\"%s\" is not a valid namespace (not a directory)\n", ns))
				// reporter.Errorf("\"%s\" is not a valid namespace (not a directory)\n", ns)
				// os.Exit(1)
			}

			findGlobs = append(findGlobs, dir)
		}
	}

	return findGlobs, searchDepth, nil
}
