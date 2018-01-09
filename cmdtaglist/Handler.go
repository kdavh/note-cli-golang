package cmdtaglist

import (
	//"fmt"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	//"github.com/kdavh/note-cli-golang/nflag"
	//"github.com/kdavh/note-cli-golang/nflow"
	//"github.com/kdavh/note-cli-golang/nparse"
	"fmt"
	parser "gopkg.in/alecthomas/kingpin.v2"
	//"os"
	//"path/filepath"
	//"regexp"
	//"strings"
)

type Handler struct {
	handler *parser.CmdClause
}

func (c *Handler) FullCommand() string {
	return c.handler.FullCommand()
}

func (c *Handler) Run(config nconfig.Config, ctx nctx.Context) bool {
	fmt.Printf("TAG COMMAND: %s\n", "list")

	return true
}

func NewHandler(app *parser.CmdClause) *Handler {
	listHandler := app.Command("ls", "Tag list.")

	return &Handler{
		handler: listHandler,
	}
}
