package nflow

import (
	"os"

	"github.com/kdavh/note-cli-golang/nconfig"
)

func ShellOpen(file string, ed nconfig.EditorInterface, rp nconfig.ReporterInterface, osCtrl *nconfig.OsCtrl) bool {
	// if editErr := ed.Open(file, cfg); editErr != nil {
	// 	rp.Error("...Error editing...")
	// 	osCtrl.Exit(1)
	// } else {
	// 	osCtrl.Exit(0)
	// }

	return true
}

func ErrExit(e error, r nconfig.ReporterInterface) bool {
	if e != nil {
		r.Debugf("ERROR, EXITING: %v", e)
		os.Exit(1)
	}
	return true
}
