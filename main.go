package main

import (
	inner "github.com/jenkins-zh/jcli-account-plugin/cmd"
	"os"
)

func main()  {
	cmd := inner.NewAccountCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
