package cmdtag

import (
	//"fmt"
	"github.com/kdavh/note-cli-golang/cmdtaglist"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nflag"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type Handler struct {
	handler     *parser.CmdClause
	namespace   *string
	listHandler *cmdtaglist.Handler
	config      *nconfig.Config
}

func (c *Handler) CanHandle(commands []string) bool {
	return len(commands) > 0 && c.handler.FullCommand() == commands[0]
}

func (c *Handler) Run(cmds []string) bool {
	if c.listHandler.CanHandle(cmds) {
		c.listHandler.Run()
	}

	return true
}

func NewHandler(app *parser.Application, config *nconfig.Config) *Handler {
	tagHandler := app.Command("tag", "Tag commands.")

	tagNamespace := nflag.HandleNamespace(tagHandler)

	tagListHandler := cmdtaglist.NewHandler(tagHandler, tagNamespace, config)

	return &Handler{
		handler:     tagHandler,
		namespace:   tagNamespace,
		listHandler: tagListHandler,
		config:      config,
	}
}
