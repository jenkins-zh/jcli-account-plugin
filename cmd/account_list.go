package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// NewAccountListCmd create a command to list all accounts
func NewAccountListCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountListCmd := accountListCmd{accountCmd}
	cmd = &cobra.Command{
		Use:   "list",
		Short: "List all jcli accounts",
		RunE:  accountListCmd.runE,
	}
	return
}

func (c *accountListCmd) runE(cmd *cobra.Command, args []string) (err error) {
	var userHome string
	if userHome, err = homedir.Dir(); err != nil {
		return
	}

	var files []os.FileInfo
	accountDir := fmt.Sprintf("%s/.jenkins-cli/data/account/", userHome)
	if files, err = ioutil.ReadDir(accountDir); err == nil {
		for _, file := range files {
			if !file.IsDir() {
				continue
			}

			fmt.Println(file.Name())
		}
	}
	return
}
