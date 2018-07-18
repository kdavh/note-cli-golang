package cmdtaglist

import (
	"strings"

	"github.com/kdavh/note-cli-golang/nconfig"
	parser "gopkg.in/alecthomas/kingpin.v2"
)

type Handler struct {
	handler   *parser.CmdClause
	namespace *string
	filter    *string
	se        nconfig.SearcherInterface
	rp        nconfig.ReporterInterface
	osCtrl    *nconfig.OsCtrl
}

func (hndl *Handler) CanHandle(commands string) bool {
	return strings.HasPrefix(commands, hndl.handler.FullCommand())
}

func (hndl *Handler) Run() bool {
	allTags, err := hndl.se.Tags(*hndl.namespace, *hndl.filter, hndl.rp)
	if err != nil {
		hndl.osCtrl.Exit(1)
	}

	hndl.rp.Reportf("\nTAGS LIST\n")
	for _, tag := range allTags {
		hndl.rp.Reportf(tag + "\n")
	}

	return true
}

func NewHandler(app *parser.CmdClause, namespace *string, se nconfig.SearcherInterface, osCtrl *nconfig.OsCtrl, rp nconfig.ReporterInterface) *Handler {
	listHandler := app.Command("ls", "Tag list.")

	tagListFilter := listHandler.Flag(
		"filter",
		"optional name filter",
	).Short('f').String()

	return &Handler{
		handler:   listHandler,
		namespace: namespace,
		filter:    tagListFilter,
		se:        se,
		rp:        rp,
		osCtrl:    osCtrl,
	}
}
