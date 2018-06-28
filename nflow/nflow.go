package nflow

import (
	"os"

	"github.com/kdavh/note-cli-golang/nconfig"
)

func ShellOpen(file string, cfg *nconfig.Config) bool {
	if editErr := cfg.Editor.Open(file, cfg); editErr != nil {
		cfg.Reporter.Error("...Error editing...")
		cfg.OsCtrl.Exit(1)
	} else {
		cfg.OsCtrl.Exit(0)
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
