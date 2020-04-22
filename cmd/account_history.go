package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// NewAccountHistoryCmd create a command to show the history of the account
func NewAccountHistoryCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountHistoryCmd := accountHistoryCmd{
		accountCmd: accountCmd,
	}
	cmd = &cobra.Command{
		Use:     "history",
		Short:   "Show the history of current account",
		PreRunE: accountHistoryCmd.preRunE,
		RunE:    accountHistoryCmd.Run,
	}
	return
}

func (c *accountHistoryCmd) preRunE(cmd *cobra.Command, args []string) (err error) {
	err = c.setName(cmd, args)
	return
}

// Run is the main point of account status command
func (c *accountHistoryCmd) Run(cmd *cobra.Command, args []string) (err error) {
	var accountDir string
	var exists bool
	if accountDir, exists, err = c.getAccountDir(); !exists {
		err = fmt.Errorf("%s is not exists", c.Name)
	} else if err != nil {
		return
	}

	var r *git.Repository
	if r, err = git.PlainOpen(accountDir); err == nil {
		var cIter object.CommitIter
		if cIter, err = r.Log(&git.LogOptions{}); err != nil {
			return
		}

		err = cIter.ForEach(func(c *object.Commit) error {
			cmd.Println(c)
			return nil
		})
	}
	return
}
