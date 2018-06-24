package nflow

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kdavh/note-cli-golang/nconfig"
)

func ShellOpen(editor string, file string, cfg *nconfig.Config) bool {
	cmd := exec.Command(editor, []string{
		"-S",
		cfg.EditorConfig,
		file,
	}...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("TAGLINE=%s", cfg.Tagline))
	cmd.Dir = cfg.NotesPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cfg.Reporter.Debugf("EDITOR COMMAND: %s %s", editor, file)

	if editErr := cmd.Run(); editErr != nil {
		cfg.Reporter.Error("...Error editing...")
		os.Exit(1)
	} else {
		os.Exit(0)
	}

	return true
}

func ErrExit(e error, reporter nconfig.ReporterInterface) bool {
	if e != nil {
		reporter.Debugf("ERROR, EXITING: %v", e)
		os.Exit(1)
	}
	return true
}
