package cmd

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
	"os"
	"strings"
)

// NewAccountCmd create a command to deal with account
func NewAccountCmd(args []string) (cmd *cobra.Command) {
	accountCmd := &accountCmd{}

	cmd = &cobra.Command {
		Use: "jcli account",
		Short: "jcli config file account",
		Long: "jcli config file account",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			accountCmd.output = cmd.OutOrStdout()
		},
		RunE: accountCmd.runE,
	}
	cmd.SetOut(os.Stdout)

	// add flags to this command
	flags := cmd.PersistentFlags();
	flags.StringVarP(&accountCmd.URL, "url", "", "",
		"The URL of a git repository")
	flags.StringVarP(&accountCmd.Username, "username", "u", "",
		"The username of the git repository")
	flags.StringVarP(&accountCmd.Password, "password", "p", "",
		"The password of the git repository")
	flags.StringVarP(&accountCmd.Name, "name", "", "",
		"Name of the account")

	sshKeyFile := fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))
	flags.StringVarP(&accountCmd.SSHKeyFile, "ssh-key-file", "", sshKeyFile,
		"SSH key file")

	// add sub-commands
	cmd.AddCommand(NewAccountAddCmd(args, accountCmd),
		NewAccountUpdateCmd(args, accountCmd),
		NewAccountSelectCmd(args, accountCmd),
		NewAccountListCmd(args, accountCmd),
		NewAccountRemoveCmd(args, accountCmd),
		NewAccountDocCmd(cmd))
	return
}

func (c *accountCmd) runE(cmd *cobra.Command, args []string) (err error) {
	var configFile string
	if configFile, err = c.getDefaultConfigPath(); err != nil {
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(configFile); err != nil {
		err = nil // it's ok if there's not config file
		cmd.Println("there's no config file found")
		return
	}

	accountConfig := accountConfig{}
	if err = yaml.Unmarshal(data, &accountConfig); err == nil {
		if accountConfig.Account == "" {
			cmd.Printf("account feature isn't enabled\n")
		} else {
			cmd.Printf("current account is %s\n", accountConfig.Account)
		}
	}
	return
}

func (c *accountCmd) getCloneOptions() (cloneOptions *git.CloneOptions) {
	cloneOptions = &git.CloneOptions{
		URL:      c.URL,
		Progress: c.output,
		Auth:     c.getAuth(),
	}
	return
}

func (c *accountCmd) getPullOptions() (pullOptions *git.PullOptions) {
	pullOptions = &git.PullOptions{
		RemoteName: "origin",
		Progress:   c.output,
		Auth:       c.getAuth(),
	}
	return
}

func (c *accountCmd) getAuth() (auth transport.AuthMethod) {
	if c.Username != "" {
		auth = &githttp.BasicAuth{
			Username: c.Username,
			Password: c.Password,
		}
	}

	if strings.HasPrefix(c.URL, "git@") {
		if sshKey, err := ioutil.ReadFile(c.SSHKeyFile); err == nil {
			signer, _ := ssh.ParsePrivateKey(sshKey)
			auth = &gitssh.PublicKeys{User: "git", Signer: signer}
		}
	}
	return
}

func (c *accountCmd) getAccountDir() (accountDir string, exists bool, err error) {
	var userHome string
	if userHome, err = homedir.Dir(); err != nil {
		return
	}

	accountDir = fmt.Sprintf("%s/.jenkins-cli/data/account/%s", userHome, c.Name)
	if _, err = os.Stat(accountDir); err == nil {
		exists = true
	} else {
		exists = false
	}
	return
}

func (c *accountCmd) getDefaultConfigPath() (configPath string, err error) {
	var userHome string
	userHome, err = homedir.Dir()
	if err == nil {
		configPath = fmt.Sprintf("%s/.jenkins-cli.yaml", userHome)
	}
	return
}

func (c* accountCmd) setName(cmd *cobra.Command, args []string) (err error) {
	if len(args) > 0 {
		c.Name = args[0]
	}

	if c.Name == "" {
		err = fmt.Errorf("name cannot be empty")
	}
	return
}
