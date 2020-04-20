package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
)

// NewAccountUpdateCmd create a command to update the account
func NewAccountUpdateCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountUpdateCmd := accountUpdateCmd{
		Reset:      true,
		accountCmd: accountCmd,
	}
	cmd = &cobra.Command{
		Use:     "update",
		Short:   "update the account",
		PreRunE: accountUpdateCmd.preRunE,
		RunE:    accountUpdateCmd.Run,
	}
	return
}

func (c *accountUpdateCmd) preRunE(cmd *cobra.Command, args []string) (err error) {
	err = c.setName(cmd, args)
	return
}

// Run is the main point of account update command
func (c *accountUpdateCmd) Run(cmd *cobra.Command, args []string) (err error) {
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

		if c.Reset {
			if err = w.Reset(&git.ResetOptions{
				Mode: git.HardReset,
			}); err != nil {
				return
			}
		}

		err = w.Pull(c.getPullOptions())
		if err == git.NoErrAlreadyUpToDate {
			err = nil // consider it's ok
		}
	}
	return
}
