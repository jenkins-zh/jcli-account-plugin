package cmd

import (
	"fmt"
	"github.com/jenkins-zh/jenkins-cli/app"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const (
	// DocTypeMarkdown represents markdown type of doc
	DocTypeMarkdown string = "Markdown"
	// DocTypeManPage represents man page type of doc
	DocTypeManPage string = "ManPage"
)

const (
	gendocFrontmatterTemplate = `---
date: %s
title: "%s"
version: %s
---
`
)

// NewAccountDocCmd can generate documents for all commands
func NewAccountDocCmd(rootCmd *cobra.Command) (cmd *cobra.Command) {
	accountDocCmd := accountDocCmd{
		RootCommand: rootCmd,
	}
	cmd = &cobra.Command{
		Use: "doc",
		Short: "Generate documents for jcli account plugin",
		RunE: accountDocCmd.runE,
	}

	flags := cmd.Flags()
	flags.StringVarP(&accountDocCmd.DocType, "doc-type", "", DocTypeMarkdown,
		"Which type of document will generate")

	err := cmd.RegisterFlagCompletionFunc("doc-type", func(cmd *cobra.Command, args []string, toComplete string) (
		i []string, directive cobra.ShellCompDirective) {
		return []string{DocTypeMarkdown, DocTypeManPage}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		cmd.PrintErrf("register flag doc-type for sub-command doc failed %#v\n", err)
	}
	return
}

func (c * accountDocCmd) runE(cmd *cobra.Command, args []string) (err error) {
	outputDir := args[0]
	if err = os.MkdirAll(outputDir, os.FileMode(0755)); err != nil {
		return
	}

	rootCmd := c.RootCommand

	switch c.DocType {
	case DocTypeMarkdown:
		now := time.Now().Format(time.RFC3339)
		prepender := func(filename string) string {
			name := filepath.Base(filename)
			base := strings.TrimSuffix(name, path.Ext(name))
			return fmt.Sprintf(gendocFrontmatterTemplate, now,
				strings.Replace(base, "_", " ", -1),
				app.GetVersion())
		}

		linkHandler := func(name string) string {
			base := strings.TrimSuffix(name, path.Ext(name))
			return "/commands/" + strings.ToLower(base) + "/"
		}

		rootCmd.DisableAutoGenTag = true
		err = doc.GenMarkdownTreeCustom(rootCmd, outputDir, prepender, linkHandler)
	case DocTypeManPage:
		header := &doc.GenManHeader{
			Title:   "Jenkins CLI account plugin",
			Section: "1",
			Source:  "Jenkins Chinese Community",
		}
		err = doc.GenManTree(rootCmd, header, outputDir)
	}

	return
}
