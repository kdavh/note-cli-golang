package test

import (
	"os/exec"
	"path/filepath"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nfs"
	"github.com/spf13/afero"
)

func CreateMockNotes(fs afero.Fs) {
	exec.Command("rm", "-r", nconfig.NotesDirMockPath()).Run()

	nfs.MkdirAll(fs, filepath.Join(nconfig.NotesDirMockPath(), "ns1"))
	nfs.MkdirAll(fs, filepath.Join(nconfig.NotesDirMockPath(), "ns2"))

	test_note, _ := fs.Create(filepath.Join(nconfig.NotesDirMockPath(), "ns1", "test_note1"))
	defer test_note.Close()
	test_note.WriteString("###-tags: tag1 tag2\n\n## Super important note")

	test_note2, _ := fs.Create(filepath.Join(nconfig.NotesDirMockPath(), "ns1", "test_note2"))
	defer test_note2.Close()
	test_note2.WriteString("###-tags: tag1 tag3\n\n## Other kind of useful note")
}

func ClearMockNotes() {
	exec.Command("rm", "-r", nconfig.NotesDirMockPath()).Run()
}
