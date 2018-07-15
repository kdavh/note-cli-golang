package neditor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"github.com/kdavh/note-cli-golang/nconfig"
)

type Editor struct {
	prog          string
	args          []string
	notesDirPath  string
	tagLineFormat string
	fs            afero.Fs
}

func (ed *Editor) Open(ns string, file string) error {
	cmdArgs := append(ed.args, file)

	cmd := exec.Command(ed.prog, cmdArgs...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("TAGLINE=%s", ed.tagLineFormat))
	cmd.Dir = filepath.Join(ed.notesDirPath, ns)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func (e *Editor) NewFile(ns string, filename string, tags []string) error {
	newNoteDir := filepath.Join(e.notesDirPath, ns)
	e.fs.MkdirAll(newNoteDir, 0755)
	filenameFull := filepath.Join(newNoteDir, filename)
	if _, statErr := e.fs.Stat(filenameFull); os.IsExist(statErr) {
		return statErr
	} else {
		file, newFileErr := e.fs.Create(filenameFull)
		if newFileErr != nil {
			return newFileErr
		}

		data := e.tagLineFormat + " " + strings.Join(tags, " ")
		fmt.Fprintf(file, data)
		file.Close()

		return nil
	}
}

func NewEditorVim(fs afero.Fs) *Editor {
	return &Editor{
		prog:          "nvim",
		args:          []string{"-S", nconfig.EditorConfigPath()},
		fs:            fs,
		notesDirPath:  nconfig.NotesDirPath(),
		tagLineFormat: nconfig.DefaultTaglineFormat(),
	}
}

func NewEditorMock(fs afero.Fs) *Editor {
	return &Editor{
		prog:          "cat",
		args:          []string{},
		fs:            fs,
		notesDirPath:  nconfig.NotesDirMockPath(),
		tagLineFormat: nconfig.DefaultTaglineFormat(),
	}
}
