package cmdnnew

import (
	"strings"
	"testing"

	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nreport"
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
	fs := afero.NewMemMapFs()
	reporter := nreport.NewMock()
	editor := nconfig.NewEditorMock(fs)
	cfg := nconfig.NewCfgMock(fs, reporter, editor)
	app := parser.New("note", "test app")

	return testHelper{
		Fs:       fs,
		Reporter: reporter,
		App:      app,
		Handler:  NewHandler(app, cfg),
	}
}

func TestNew(t *testing.T) {
	h := setupTest()
	strings.Split(parser.MustParse(h.App.Parse(strings.Split("new test-file.md -tt1,t2", " "))), " ")
	h.Handler.Run()

	stat, err := h.Fs.Stat("notes/test-file.md")
	assert.NotEmpty(t, stat)
	assert.Nil(t, err)
}

func TestNewFail(t *testing.T) {
	h := setupTest()
	strings.Split(parser.MustParse(h.App.Parse(strings.Split("new bad-test-file -tt1,t2", " "))), " ")

	defer func() {
		r := recover()
		assert.Equal(t, 1, r, "exited with status 1 because filename not correct")
		assert.Equal(t,
			[]string{"bad-test-file must end with `.md`, exiting\n"},
			h.Reporter.ErrorCalls,
		)

		stat, err := h.Fs.Stat("notes/bad-test-file")
		assert.Empty(t, stat)
		assert.NotNil(t, err)

	}()
	h.Handler.Run()
}
