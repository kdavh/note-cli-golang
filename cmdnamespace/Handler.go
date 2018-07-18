package cmdnamespace

import (
	"strings"

	"github.com/kdavh/note-cli-golang/cmdnamespacelist"
	"github.com/kdavh/note-cli-golang/nconfig"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type Handler struct {
	handler     *parser.CmdClause
	listHandler *cmdnamespacelist.Handler
}

func (hndl *Handler) CanHandle(commands string) bool {
	return strings.HasPrefix(commands, hndl.handler.FullCommand())
}

func (c *Handler) Run(cmds string) bool {
	if c.listHandler.CanHandle(cmds) {
		c.listHandler.Run()
	}

	return true
}

func NewHandler(app *parser.Application, se nconfig.SearcherInterface, osCtrl *nconfig.OsCtrl, rp nconfig.ReporterInterface) *Handler {
	namespaceHandler := app.Command("namespace", "Namespace commands.")

	namespaceListHandler := cmdnamespacelist.NewHandler(namespaceHandler, se, osCtrl, rp)

	return &Handler{
		handler:     namespaceHandler,
		listHandler: namespaceListHandler,
	}
}
