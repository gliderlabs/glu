package main

import (
	"fmt"
	"os"

	"github.com/progrium/go-shell"
	"github.com/spf13/cobra"
)

func init() {
	Glu.AddCommand(circleciCmd)
}

var circleciCmd = &cobra.Command{
	Use:   "circleci",
	Short: "Sets up a circleci environment",
	Run: func(cmd *cobra.Command, args []string) {
		defer shell.ErrExit()
		shell.Trace = true
		sh("rm -f ~/.gitconfig")
		if exists("VERSION") && os.Getenv("CIRCLE_BRANCH") != "release" {
			build := fmt.Sprintf("build-%s", os.Getenv("CIRCLE_BUILD_NUM"))
			sh("echo", q(build), "> VERSION")
		}
	},
}
