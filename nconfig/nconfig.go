package nconfig

import (
	"os"

	"github.com/spf13/afero"
)

type OsCtrl struct {
	Exit    func(int)
	IsExist func(err error) bool
}

type SearcherInterface interface {
	Notes(string, string, string, *Config) ([]string, error)
}

type ReporterInterface interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Error(string)
}

type EditorInterface interface {
	Open(file string, cfg *Config) error
}

func NewOsCtrl() OsCtrl {
	return OsCtrl{
		Exit:    os.Exit,
		IsExist: os.IsExist,
	}
}

type Config struct {
	Searcher  SearcherInterface
	Tagline   string
	NotesPath string
	Fs        afero.Fs
	OsCtrl    OsCtrl
	Reporter  ReporterInterface
	Editor    EditorInterface
}
