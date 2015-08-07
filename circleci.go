package main

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	GluCmd.AddCommand(circleciCmd)
}

var circleciCmd = &cobra.Command{
	Use:   "circleci",
	Short: "Sets up a circleci environment",
	Run: func(cmd *cobra.Command, args []string) {
		sh("rm -f ~/.gitconfig")
		if exists("VERSION") && os.Getenv("CIRCLE_BRANCH") != "release" {
			writeFile("VERSION", "build-"+os.Getenv("CIRCLE_BUILD_NUM"))
		}
	},
}
