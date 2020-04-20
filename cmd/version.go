package cmd

import (
	"github.com/jenkins-zh/jenkins-cli/app"
	"github.com/spf13/cobra"
)

// NewVersionCmd create a command to show the version
func NewVersionCmd() (cmd *cobra.Command) {
	versionCmd := versionCmd{}
	cmd = &cobra.Command{
		Use:     "version",
		Short:   "Show the version of this plugin",
		RunE:    versionCmd.RunE,
	}
	return
}

// RunE is the main entry point of this command
func (c *versionCmd) RunE(cmd *cobra.Command, args []string) (err error) {
	cmd.Printf("Version: %s\n", app.GetVersion())
	cmd.Printf("Last Commit: %s\n", app.GetCommit())
	//cmd.Printf("Build Date: %s\n", app.GetDate())
	return
}
