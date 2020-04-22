package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

// NewAccountAddCmd create a command to add an account
func NewAccountAddCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountAddCmd := accountAddCmd{accountCmd}
	cmd = &cobra.Command{
		Use:     "add",
		Short:   "Add a account for jcli config",
		PreRunE: accountAddCmd.preRunE,
		RunE:    accountAddCmd.Run,
	}
	return
}

func (c *accountAddCmd) preRunE(cmd *cobra.Command, args []string) (err error) {
	err = c.setName(cmd, args)
	return
}

// Run is the main point of account add command
func (c *accountAddCmd) Run(cmd *cobra.Command, args []string) (err error) {
	var accountDir string
	var exists bool
	if accountDir, exists, err = c.getAccountDir(); exists {
		err = fmt.Errorf("%s is exists", c.Name)
	} else if !os.IsNotExist(err) {
		return
	}

	c.installProtocol()
	if _, err = git.PlainOpen(accountDir); err == nil {
		err = fmt.Errorf("%s is exists", c.Name)
		return
	}
	_, err = git.PlainClone(accountDir, false, c.getCloneOptions())
	return
}
