package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
)

// NewAccountStatusCmd create a command to show the status of the account
func NewAccountStatusCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountStatusCmd := accountStatusCmd{
		accountCmd: accountCmd,
	}
	cmd = &cobra.Command{
		Use:     "status",
		Short:   "Show the status of current account",
		PreRunE: accountStatusCmd.preRunE,
		RunE:    accountStatusCmd.Run,
	}
	return
}

func (c *accountStatusCmd) preRunE(cmd *cobra.Command, args []string) (err error) {
	err = c.setName(cmd, args)
	return
}

// Run is the main point of account status command
func (c *accountStatusCmd) Run(cmd *cobra.Command, args []string) (err error) {
	var accountDir string
	var exists bool
	if accountDir, exists, err = c.getAccountDir(); !exists {
		err = fmt.Errorf("%s is not exists", c.Name)
	} else if err != nil {
		return
	}

	var r *git.Repository
	if r, err = git.PlainOpen(accountDir); err == nil {
		var w *git.Worktree
		if w, err = r.Worktree(); err != nil {
			return
		}

		var status git.Status
		if status, err = w.Status(); err == nil {
			cmd.Print(status.String())
		}
	}
	return
}
