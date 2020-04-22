package cmd

import (
	jcli "github.com/jenkins-zh/jenkins-cli/app/config"
	"github.com/spf13/cobra"
	"io"
)

type (
	accountCmd struct {
		URL        string
		Username   string
		Password   string
		SSHKeyFile string

		Name string

		Proxy     string
		ProxyAuth string

		output io.Writer
	}

	accountAddCmd struct {
		*accountCmd
	}

	accountSelectCmd struct {
		*accountCmd
	}

	accountUpdateCmd struct {
		Reset bool
		*accountCmd
	}

	accountListCmd struct {
		*accountCmd
	}

	accountStatusCmd struct {
		*accountCmd
	}

	accountCommitCmd struct {
		*accountCmd
		Message string
	}

	accountRemoveCmd struct {
		*accountCmd
	}

	accountHistoryCmd struct {
		*accountCmd
	}

	accountConfig struct {
		jcli.Config `yaml:",inline"`
		Account     string `yaml:"account"`
	}

	accountDocCmd struct {
		RootCommand *cobra.Command
		DocType     string
	}
)

type (
	versionCmd struct {
	}
)
