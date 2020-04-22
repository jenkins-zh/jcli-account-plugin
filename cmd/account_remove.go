package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
)

// NewAccountRemoveCmd create a command to remove a account
func NewAccountRemoveCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountRemoveCmd := accountRemoveCmd{accountCmd}
	cmd = &cobra.Command{
		Use:     "remove",
		Short:   "Remove a jcli account",
		PreRunE: accountRemoveCmd.preRunE,
		RunE:    accountRemoveCmd.runE,
	}
	return
}

func (c *accountRemoveCmd) preRunE(cmd *cobra.Command, args []string) (err error) {
	err = c.setName(cmd, args)
	return
}

func (c *accountRemoveCmd) runE(cmd *cobra.Command, args []string) (err error) {
	var userHome string
	if userHome, err = homedir.Dir(); err != nil {
		return
	}

	accountDir := fmt.Sprintf("%s/.jenkins-cli/data/account/%s", userHome, c.Name)
	err = os.RemoveAll(accountDir)
	return
}
