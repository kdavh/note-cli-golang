package nsearch

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kdavh/note-cli-golang/cmdparse"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nparse"
)

type searcher struct {
	prog string
}

func (s *searcher) Notes(namespace string, tagsQuery string, textQuery string, cfg *nconfig.Config) ([]string, error) {
	fileGlobs, searchDepth := cmdparse.FileGlobs(namespace, cfg)

	var tagsLookaheads []string
	for _, tag := range nparse.CommaSplit(tagsQuery) {
		tagsLookaheads = append(tagsLookaheads, fmt.Sprintf("(?=.*\\s+%s(\\s+|\\$))", tag))
	}

	cmd := exec.Command(s.prog, append([]string{
		cfg.Tagline + strings.Join(tagsLookaheads, "|"),
		"--files-with-matches",
		"--depth=" + searchDepth,
	}, fileGlobs...)...)

	cfg.Reporter.Debugf("SEARCH COMMAND: %s\n", strings.Join(cmd.Args, " "))

	if output, cmdErr := cmd.Output(); cmdErr != nil {
		return []string{}, cmdErr
	} else {
		return strings.Split(strings.TrimSpace(string(output)), "\n"), nil
	}
}
