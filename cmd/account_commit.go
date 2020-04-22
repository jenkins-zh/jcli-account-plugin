package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// NewAccountCommitCmd create a command to commit the changes
func NewAccountCommitCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountCommitCmd := accountCommitCmd{
		accountCmd: accountCmd,
	}
	cmd = &cobra.Command{
		Use:     "commit",
		Short:   "Commit the changes of current account",
		Long: `Commit and push the changes of current account`,
		PreRunE: accountCommitCmd.preRunE,
		RunE:    accountCommitCmd.Run,
	}

	flags := cmd.Flags()
	flags.StringVarP(&accountCommitCmd.Message, "message", "m", "",
		`The commit message of account changes`)
	return
}

func (c *accountCommitCmd) preRunE(cmd *cobra.Command, args []string) (err error) {
	if err = c.setName(cmd, args); err != nil {
		return
	}
	if c.Message == "" {
		err = fmt.Errorf("message is empty")
	}
	return
}

// Run is the main point of account status command
func (c *accountCommitCmd) Run(cmd *cobra.Command, args []string) (err error) {
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

		if _, err = w.Add("."); err != nil {
			return
		}

		opts := &git.CommitOptions{
			All: true,
			Author: &object.Signature{
				Name: "jcli",
				Email: "jcli@github.com",
			},
		}
		if _, err = w.Commit(c.Message, opts); err == nil {
			err = r.Push(c.getPushOptions())
		}
	}
	return
}
