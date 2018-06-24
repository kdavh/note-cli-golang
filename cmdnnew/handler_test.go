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

func TestNewHandlerRun(t *testing.T) {
	assert := assert.New(t)
	fs := afero.NewMemMapFs()
	reporter := nreport.NewMock()
	cfg := nconfig.NewCfgMock(fs, reporter)

	app := parser.New("note", "test app")

	handler := NewHandler(app, cfg)

	strings.Split(parser.MustParse(app.Parse(strings.Split("new test-file.md -tt1,t2", " "))), " ")
	handler.Run()

	stat, err := fs.Stat("notes/test-file.md")
	assert.NotEmpty(stat)
	assert.Nil(err)
}
