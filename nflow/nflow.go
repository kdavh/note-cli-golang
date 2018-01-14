package nflow

import (
	"fmt"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nlog"
	"os"
	"os/exec"
)

func ShellOpen(editor string, file string, logger *nlog.Logger, cfg *nconfig.Config) bool {
	cmd := exec.Command(editor, []string{
		"-S",
		cfg.EditorConfig,
		file,
	}...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("TAGLINE=%s", cfg.Tagline))
	cmd.Dir = cfg.NotesPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	logger.Debugf("EDITOR COMMAND: %s %s", editor, file)

	if editErr := cmd.Run(); editErr != nil {
		logger.Error("...Error editing...")
		os.Exit(1)
	} else {
		os.Exit(0)
	}

	return true
}

func ErrExit(e error, logger *nlog.Logger) bool {
	if e != nil {
		logger.Debugf("ERROR, EXITING: %v", e)
		os.Exit(1)
	}
	return true
}
