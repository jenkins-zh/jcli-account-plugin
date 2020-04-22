package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

// NewAccountSelectCmd create a command to select a account
func NewAccountSelectCmd(args []string, accountCmd *accountCmd) (cmd *cobra.Command) {
	accountSelectCmd := accountSelectCmd{accountCmd}
	cmd = &cobra.Command{
		Use:     "select",
		Short:   "Select a jcli config account as the default one",
		PreRunE: accountSelectCmd.PreRunE,
		RunE:    accountSelectCmd.Run,
	}
	return
}

// PreRunE do the options checking
func (c *accountSelectCmd) PreRunE(cmd *cobra.Command, args []string) (err error) {
	err = c.setName(cmd, args)
	return
}

// Run is the main point of account select command
func (c *accountSelectCmd) Run(cmd *cobra.Command, args []string) (err error) {
	var configFile string
	if configFile, err = c.getDefaultConfigPath(); err != nil {
		return
	}

	exist := true
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		exist = false
		if err = ioutil.WriteFile(configFile, []byte{}, 0664); err != nil {
			return
		}
	}

	var data []byte
	if data, err = ioutil.ReadFile(configFile); err == nil {
		config := &accountConfig{}
		if err = yaml.Unmarshal(data, config); err == nil {
			if exist && config.Account == "" {
				var userHome string
				if userHome, err = homedir.Dir(); err != nil {
					return
				}

				backupFile := fmt.Sprintf("%s/.jenkins-cli/data/account/.jenkins-cli.yaml", userHome)

				cmd.Printf("current config file don't support account feature, will move it to %s\n", backupFile)

				if err = ioutil.WriteFile(backupFile, data, 0664); err != nil {
					return
				}
			}

			if data, err = c.getAccountConfig(); err != nil {
				return
			}

			// let the target config as default
			if err = ioutil.WriteFile(configFile, data, 0664); err != nil {
				return
			}
		}
	}
	return
}

func (c *accountSelectCmd) getAccountConfig() (data []byte, err error) {
	// read target config file
	accountDir, _, _ := c.getAccountDir()
	accountConfigFile := path.Join(accountDir, ".jenkins-cli.yaml")
	if data, err = ioutil.ReadFile(accountConfigFile); err != nil {
		return
	}

	config := &accountConfig{}
	if err = yaml.Unmarshal(data, &config); err != nil {
		return
	}

	// make sure the account field is correct
	config.Account = c.Name
	if data, err = yaml.Marshal(&config); err != nil {
		return
	}
	err = ioutil.WriteFile(accountConfigFile, data, 0664)
	return
}
