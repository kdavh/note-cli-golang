package cmdparse

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nreport"
	"github.com/spf13/afero"
)

func TestFileGlobs(t *testing.T) {
	assert := assert.New(t)
	fs := afero.NewMemMapFs()
	reporter := nreport.NewMock()
	editor := nconfig.NewEditorMock(fs)
	cfg := nconfig.NewCfgMock(fs, reporter, editor)

	createMockNotes(fs)
	globs, searchDepth := FileGlobs("namespace1", cfg)
	assert.Equal([]string{"notes/namespace1"}, globs)
	assert.Equal("0", searchDepth)
	assert.Equal([]string{}, reporter.ErrorCalls)
	reporter.Reset()

	globs, searchDepth = FileGlobs("namespace1,ns2", cfg)
	assert.Equal([]string{"notes/namespace1", "notes/ns2"}, globs)
	assert.Equal("0", searchDepth)
	assert.Equal([]string{}, reporter.ErrorCalls)
	reporter.Reset()

	globs, searchDepth = FileGlobs("", cfg)
	assert.Equal([]string{"notes"}, globs)
	assert.Equal("0", searchDepth)
	assert.Equal([]string{}, reporter.ErrorCalls)
	reporter.Reset()

	globs, searchDepth = FileGlobs("*", cfg)
	assert.Equal([]string{"notes"}, globs)
	assert.Equal("1", searchDepth)
	assert.Equal([]string{}, reporter.ErrorCalls)
	reporter.Reset()

	defer func() {
		r := recover()
		assert.Equal(1, r, "exited with status 1 because namespace doesn't exist")
		assert.Equal([]string{"open notes/nsDoesNotExist: file does not exist\n"}, reporter.ErrorCalls)
	}()
	globs, searchDepth = FileGlobs("nsDoesNotExist", cfg)
}

func createMockNotes(fs afero.Fs) {
	fs.Mkdir("notes", os.ModeDir)
	fs.Mkdir("notes/namespace1", os.ModeDir)
	fs.Mkdir("notes/ns2", os.ModeDir)
	test_note, _ := fs.Create("notes/namespace1/test_note")
	defer test_note.Close()
	test_note.WriteString("###-tags: tag1, tag2\n\n## Super important note")
}
