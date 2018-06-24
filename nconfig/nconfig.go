package nconfig

import (
	"os"

	"github.com/spf13/afero"
)

type exitFunc func(int)

type OsCtrl struct {
	Exit    exitFunc
	IsExist func(err error) bool
}

type ReporterInterface interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Error(string)
}

func NewOsCtrl() OsCtrl {
	return OsCtrl{
		Exit:    os.Exit,
		IsExist: os.IsExist,
	}
}

type Config struct {
	SearchApp    string
	Editor       string
	EditorConfig string
	Tagline      string
	NotesPath    string
	Fs           afero.Fs
	OsCtrl       OsCtrl
	Reporter     ReporterInterface
}
