package cmdparse

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nlog"
	"github.com/spf13/afero"
)

func TestFileGlobs(t *testing.T) {
	assert := assert.New(t)

	c := &nconfig.Config{
		SearchApp:    "ag",
		Editor:       "nvim",
		EditorConfig: filepath.Join("mocks", "editor-config"),
		Tagline:      "###-tags-:",
		NotesPath:    "notes",
		Fs:           afero.NewMemMapFs(),
		OsCtrl: nconfig.OsCtrl{
			Exit: func(code int) {},
		},
	}

	ctx := &nctx.Context{
		nlog.New(nlog.ERROR),
	}

	createMockNotes(c.Fs)
	globs, searchDepth := FileGlobs("namespace1", c, ctx)
	assert.Equal([]string{"notes/namespace1"}, globs)
	assert.Equal("0", searchDepth)

	globs, searchDepth = FileGlobs("namespace1,ns2", c, ctx)
	assert.Equal([]string{"notes/namespace1", "notes/ns2"}, globs)
	assert.Equal("0", searchDepth)

	globs, searchDepth = FileGlobs("", c, ctx)
	assert.Equal([]string{"notes"}, globs)
	assert.Equal("0", searchDepth)

	globs, searchDepth = FileGlobs("*", c, ctx)
	assert.Equal([]string{"notes"}, globs)
	assert.Equal("1", searchDepth)
}

func createMockNotes(fs afero.Fs) {
	fs.Mkdir("notes", os.ModeDir)
	fs.Mkdir("notes/namespace1", os.ModeDir)
	test_note, _ := fs.Create("notes/namespace1/test_note")
	defer test_note.Close()
	test_note.WriteString("###-tags: tag1, tag2\n\n## Super important note")
}
