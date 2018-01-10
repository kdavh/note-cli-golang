package cmdtaglist

import (
	"fmt"
	"github.com/kdavh/note-cli-golang/cmdhelp"
	"github.com/kdavh/note-cli-golang/nconfig"
	"github.com/kdavh/note-cli-golang/nctx"
	parser "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

type Handler struct {
	handler   *parser.CmdClause
	namespace *string
	config    *nconfig.Config
	ctx       *nctx.Context
}

func (c *Handler) CanHandle(commands []string) bool {
	handlerCmds := strings.Split(c.handler.FullCommand(), " ")

	return len(commands) >= len(handlerCmds) &&
		handlerCmds[0] == commands[0] &&
		handlerCmds[1] == commands[1]
}

func (c *Handler) Run() bool {
	//ctx.logger.out("TAGS:")
	//notes_glob = "#{cfg.notes_dir}/*.md"
	//namespaced_notes_glob = "#{cfg.notes_dir}/**/*.md"
	//search_prg = 'ag'

	//find_tag_lines_cmd = "#{search_prg} --nofilename \"#{cfg.file_tags_header}\" #{notes_glob} #{namespaced_notes_glob}"

	//ctx.logger.debug find_tag_lines_cmd

	//tag_lines = `#{find_tag_lines_cmd}`.split("\n").map { |tag_line| tag_line.sub(cfg.file_tags_header, '').strip }

	//tags = SortedSet.new

	//tag_lines.each do |line|
	//tags.merge(
	//line.sub(cfg.file_tags_header, '').split(/\s*,\s*/)
	//)
	//end

	//ctx.logger.out "\t" + tags.to_a.join("\n\t")
	cfg := c.config
	ctx := c.ctx

	findGlobs, searchDepth := cmdhelp.FileGlobs(*c.namespace, cfg, ctx)
	tagFindCmd := exec.Command(cfg.SearchApp, append([]string{"--nofilename", cfg.Tagline, "--depth=" + searchDepth}, findGlobs...)...)

	if output, cmdErr := tagFindCmd.Output(); cmdErr != nil {
		ctx.Logger.Errorf("COMMAND FAILED: %s\n\nERROR: %s", tagFindCmd, cmdErr)

		os.Exit(1)
	} else {
		allTagsMap := make(map[string]bool)
		var allTags []string
		lines := regexp.MustCompile(`\n+`).Split(
			strings.TrimSpace(string(output)), -1,
		)

		for _, line := range lines {
			tags := regexp.MustCompile(`\s+`).Split(
				strings.TrimSpace(strings.Replace(line, cfg.Tagline, "", 1)), -1,
			)

			for _, tag := range tags {
				allTagsMap[tag] = true
			}
		}

		for tag, _ := range allTagsMap {
			allTags = append(allTags, tag)
		}

		fmt.Println("\nTAGS LIST")
		sort.Strings(allTags)
		for _, tag := range allTags {
			fmt.Println(tag)
		}
	}

	return true
}

func NewHandler(app *parser.CmdClause, namespace *string, config *nconfig.Config, ctx *nctx.Context) *Handler {
	listHandler := app.Command("ls", "Tag list.")

	return &Handler{
		handler:   listHandler,
		namespace: namespace,
		config:    config,
		ctx:       ctx,
	}
}
