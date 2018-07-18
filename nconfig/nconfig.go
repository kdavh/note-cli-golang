package nconfig

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

var devDir string = filepath.Join(os.Getenv("HOME"), "dev")

func DefaultTaglineFormat() string {
	return "###-tags-:"
}
func NotesDirPath() string {
	return filepath.Join(devDir, "note-app-notes", "notes")
}
func NotesDirMockPath() string {
	return "/tmp/note-cli-golang-test"
}
func EditorConfigPath() string {
	return filepath.Join(devDir, "note-app-vim", "vim-note-config.vimrc")
}

type OsCtrl struct {
	Exit func(int)
}

type SearcherInterface interface {
	Notes(string, []string, string, ReporterInterface) ([]string, error)
	Tags(string, string, ReporterInterface) ([]string, error)
}

type ReporterInterface interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Error(string)
	Reportf(string, ...interface{})
	Prompt() string
}

type EditorInterface interface {
	Open(file string, rp ReporterInterface) error
	NewFile(namespace string, filename string, tags []string) error
}

func NewOsCtrl() *OsCtrl {
	return &OsCtrl{
		Exit: os.Exit,
	}
}

type Config struct {
	Searcher  SearcherInterface
	Tagline   string
	NotesPath string
	Fs        afero.Fs
	OsCtrl    *OsCtrl
	Reporter  ReporterInterface
	Editor    EditorInterface
}
