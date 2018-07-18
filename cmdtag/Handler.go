package cmdtag

import (
	"strings"

	"github.com/kdavh/note-cli-golang/cmdtaglist"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nflag"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type Handler struct {
	handler     *parser.CmdClause
	namespace   *string
	listHandler *cmdtaglist.Handler
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
	tagHandler := app.Command("tag", "Tag commands.")

	tagNamespace := nflag.HandleNamespace(tagHandler)

	tagListHandler := cmdtaglist.NewHandler(tagHandler, tagNamespace, se, osCtrl, rp)

	return &Handler{
		handler:     tagHandler,
		namespace:   tagNamespace,
		listHandler: tagListHandler,
	}
}
