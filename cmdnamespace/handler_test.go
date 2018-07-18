package cmdnamespace

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
	cmds := parser.MustParse(h.App.Parse(strings.Split("namespace ls", " ")))
	hndl.Run(cmds)
	assert.Contains(t, h.Rp.ReportCalls[0], "NAMESPACES LIST")
	assert.Contains(t, h.Rp.ReportCalls[1], "/")
	assert.Contains(t, h.Rp.ReportCalls[2], "ns1/")
	assert.Contains(t, h.Rp.ReportCalls[3], "ns2/")
}
