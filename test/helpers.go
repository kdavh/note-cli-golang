package test

import (
	"os/exec"
	"path/filepath"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/neditor"
	"github.com/kdavh/note-cli-golang/nreport"
	"github.com/kdavh/note-cli-golang/nsearch"
	"github.com/spf13/afero"
	parser "gopkg.in/alecthomas/kingpin.v2"

	"github.com/kdavh/note-cli-golang/nfs"
)

type TestHelper struct {
	Fs     afero.Fs
	Rp     *nreport.ReporterMock
	App    *parser.Application
	OsCtrl *nconfig.OsCtrl
	Ed     *neditor.Editor
	Se     nconfig.SearcherInterface
}

func SetupTest() TestHelper {
	fs := afero.NewOsFs()
	reporter := nreport.NewMock()
	osCtrl := nconfig.NewOsCtrlMock()
	editor := neditor.NewEditorMock(fs)
	searcher := nsearch.NewSearcherMock(fs)
	app := parser.New("note", "test app")

	CreateMockNotes(fs)

	return TestHelper{
		Fs:     fs,
		Rp:     reporter,
		App:    app,
		OsCtrl: osCtrl,
		Ed:     editor,
		Se:     searcher,
	}
}

func CreateMockNotes(fs afero.Fs) {
	exec.Command("rm", "-r", nconfig.NotesDirMockPath()).Run()

	nfs.MkdirAll(fs, filepath.Join(nconfig.NotesDirMockPath(), "ns1"))
	nfs.MkdirAll(fs, filepath.Join(nconfig.NotesDirMockPath(), "ns2"))

	test_note, _ := fs.Create(filepath.Join(nconfig.NotesDirMockPath(), "ns1", "test_note1"))
	defer test_note.Close()
	test_note.WriteString("###-tags-: tag1 tag2\n\n## Super important note")

	test_note2, _ := fs.Create(filepath.Join(nconfig.NotesDirMockPath(), "ns1", "test_note2"))
	defer test_note2.Close()
	test_note2.WriteString("###-tags-: tag1 tag3\n\n## Other kind of useful note")
}

func ClearMockNotes() {
	exec.Command("rm", "-r", nconfig.NotesDirMockPath()).Run()
}
