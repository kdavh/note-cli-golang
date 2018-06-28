package nconfig

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/afero"
)

type OsCtrl struct {
	Exit    func(int)
	IsExist func(err error) bool
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

type editor struct {
	editorProg   string
	editorConfig string
}

func (e *editor) Open(file string, cfg *Config) error {
	cfg.Reporter.Debugf("EDITOR COMMAND: %s %s", e.editorProg, file)

	cmd := exec.Command(e.editorProg, []string{
		"-S",
		e.editorConfig,
		file,
	}...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("TAGLINE=%s", cfg.Tagline))
	cmd.Dir = cfg.NotesPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func NewOsCtrl() OsCtrl {
	return OsCtrl{
		Exit:    os.Exit,
		IsExist: os.IsExist,
	}
}

func NewEditorVim(configRoot string) *editor {
	return &editor{
		editorProg:   "nvim",
		editorConfig: filepath.Join(configRoot, "note-app-vim", "vim-note-config.vimrc"),
	}
}

type Config struct {
	SearchApp string
	Tagline   string
	NotesPath string
	Fs        afero.Fs
	OsCtrl    OsCtrl
	Reporter  ReporterInterface
	Editor    EditorInterface
}
