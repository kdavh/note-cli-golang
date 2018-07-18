package cmdtag

import (
	"strings"
	"testing"

	"github.com/kdavh/note-cli-golang/test"
	"github.com/stretchr/testify/assert"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

func setupTest() (test.TestHelper, *Handler) {
	h := test.SetupTest()
	return h, NewHandler(h.App, h.Se, h.OsCtrl, h.Rp)
}

func TestList(t *testing.T) {
	h, hndl := setupTest()
	cmds := parser.MustParse(h.App.Parse(strings.Split("tag -nns1 ls", " ")))
	hndl.Run(cmds)
	assert.Contains(t, h.Rp.ReportCalls[0], "TAGS LIST")
	assert.Contains(t, h.Rp.ReportCalls[1], "tag1")
	assert.Contains(t, h.Rp.ReportCalls[2], "tag2")
	assert.Contains(t, h.Rp.ReportCalls[3], "tag3")
}

func TestListFilteredPrefix(t *testing.T) {
	h, hndl := setupTest()
	cmds := parser.MustParse(h.App.Parse(strings.Split("tag -nns1 ls -ftag1", " ")))
	hndl.Run(cmds)
	assert.Len(t, h.Rp.ReportCalls, 2)
	assert.Contains(t, h.Rp.ReportCalls[0], "TAGS LIST")
	assert.Contains(t, h.Rp.ReportCalls[1], "tag1")
}

func TestListFilteredRegex(t *testing.T) {
	h, hndl := setupTest()
	cmds := parser.MustParse(h.App.Parse(strings.Split("tag -nns1 ls -f ^.ag1", " ")))
	hndl.Run(cmds)
	assert.Len(t, h.Rp.ReportCalls, 2)
	assert.Contains(t, h.Rp.ReportCalls[0], "TAGS LIST")
	assert.Contains(t, h.Rp.ReportCalls[1], "tag1")
}

func TestListFilteredRegexInvalid(t *testing.T) {
	h, hndl := setupTest()
	cmds := parser.MustParse(h.App.Parse(strings.Split("tag -nns1 ls -f ^..ag1", " ")))
	defer func() {
		r := recover()
		assert.Equal(t, 1, r)
	}()
	hndl.Run(cmds)
}
