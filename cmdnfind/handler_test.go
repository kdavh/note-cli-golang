package cmdnfind

import (
	"strings"
	"testing"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/neditor"
	"github.com/kdavh/note-cli-golang/nreport"
	"github.com/kdavh/note-cli-golang/nsearch"
	"github.com/kdavh/note-cli-golang/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type testHelper struct {
	Fs       afero.Fs
	Reporter *nreport.ReporterMock
	App      *parser.Application
	Handler  *Handler
}

func setupTest() testHelper {
	fs := afero.NewOsFs()
	reporter := nreport.NewMock()
	osCtrl := nconfig.NewOsCtrlMock()
	editor := neditor.NewEditorMock(fs)
	searcher := nsearch.NewSearcherMock(fs)
	app := parser.New("note", "test app")

	test.CreateMockNotes(fs)

	return testHelper{
		Fs:       fs,
		Reporter: reporter,
		App:      app,
		Handler:  NewHandler(app, searcher, editor, osCtrl, reporter),
	}
}

func TestFind(t *testing.T) {
	h := setupTest()
	parser.MustParse(h.App.Parse(strings.Split("find -ttag1,tag2 -nns1", " ")))
	h.Handler.Run()
	assert.Contains(t,
		h.Reporter.ReportCalls[0],
		"FOUND:",
	)
	assert.Contains(t,
		h.Reporter.ReportCalls[1],
		"/tmp/note-cli-golang-test/ns1/test_note1",
	)
}

func TestFindWrongNamespace(t *testing.T) {
	h := setupTest()
	parser.MustParse(h.App.Parse(strings.Split("find -ttag1,tag2 -nns2", " ")))

	defer func() {
		r := recover()
		assert.Equal(t, 1, r, "did not exit with status 1 as expected")
		assert.Equal(t,
			[]string{NO_NOTES_FOUND + "\n"},
			h.Reporter.ErrorCalls,
		)
	}()
	h.Handler.Run()
}

func TestFindWrongTag(t *testing.T) {
	h := setupTest()
	parser.MustParse(h.App.Parse(strings.Split("find -tnontag1,nontag2 -nns1", " ")))

	defer func() {
		r := recover()
		assert.Equal(t, 1, r, "did not exit with status 1 as expected")
		assert.Equal(t,
			[]string{NO_NOTES_FOUND + "\n"},
			h.Reporter.ErrorCalls,
		)
	}()
	h.Handler.Run()
}
