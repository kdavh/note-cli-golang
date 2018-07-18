package cmdnamespacelist

import (
	"strings"

	"github.com/kdavh/note-cli-golang/nconfig"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type Handler struct {
	handler *parser.CmdClause
	filter  *string
	se      nconfig.SearcherInterface
	rp      nconfig.ReporterInterface
	osCtrl  *nconfig.OsCtrl
}

func (hndl *Handler) CanHandle(commands string) bool {
	return strings.HasPrefix(commands, hndl.handler.FullCommand())
}

func (hndl *Handler) Run() bool {
	allNamespaces, err := hndl.se.Namespaces()
	if err != nil {
		hndl.osCtrl.Exit(1)
	}

	hndl.rp.Reportf("\nNAMESPACES LIST\n")
	for _, namespace := range allNamespaces {
		hndl.rp.Reportf(namespace + "\n")
	}

	return true
}

func NewHandler(app *parser.CmdClause, se nconfig.SearcherInterface, osCtrl *nconfig.OsCtrl, rp nconfig.ReporterInterface) *Handler {
	listHandler := app.Command("ls", "Tag list.")

	namespaceListFilter := listHandler.Flag(
		"filter",
		"optional name filter",
	).Short('f').String()

	return &Handler{
		handler: listHandler,
		filter:  namespaceListFilter,
		se:      se,
		rp:      rp,
		osCtrl:  osCtrl,
	}
}
