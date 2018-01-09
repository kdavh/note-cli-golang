package cmdtag

import (
	//"fmt"
	"github.com/kdavh/note-cli-golang/cmdtaglist"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	"github.com/kdavh/note-cli-golang/nflag"
	//"github.com/kdavh/note-cli-golang/nflow"
	//"github.com/kdavh/note-cli-golang/nparse"
	parser "gopkg.in/alecthomas/kingpin.v2"
	//"os"
	//"path/filepath"
	//"regexp"
	"strings"
)

type Handler struct {
	handler     *parser.CmdClause
	namespace   *string
	listHandler *cmdtaglist.Handler
}

func (c *Handler) FullCommand() string {
	return c.handler.FullCommand()
}

func (c *Handler) Run(config nconfig.Config, ctx nctx.Context, cmds []string) bool {
	switch strings.Join(cmds, " ") {
	case c.listHandler.FullCommand():
		c.listHandler.Run(config, ctx)
	}

	return true
}

func NewHandler(app *parser.Application) *Handler {
	tagHandler := app.Command("tag", "Tag commands.")

	tagNamespace := nflag.HandleNamespace(tagHandler)

	tagListHandler := cmdtaglist.NewHandler(tagHandler)

	return &Handler{
		handler:     tagHandler,
		namespace:   tagNamespace,
		listHandler: tagListHandler,
	}
}
