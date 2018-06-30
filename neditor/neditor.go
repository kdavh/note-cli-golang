package neditor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kdavh/note-cli-golang/nconfig"
)

type editor struct {
	editorProg   string
	editorConfig string
}

func (e *editor) Open(file string, cfg *nconfig.Config) error {
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

func NewEditorVim(configRoot string) *editor {
	return &editor{
		editorProg:   "nvim",
		editorConfig: filepath.Join(configRoot, "note-app-vim", "vim-note-config.vimrc"),
	}
}
