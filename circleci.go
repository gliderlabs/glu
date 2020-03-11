package main

import (
	"fmt"
	"os"
	"path/filepath"

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
		fmt.Printf("glu/%s\n\n", Version)
		defer shell.ErrExit()
		shell.Trace = true
		sh("rm -f ~/.gitconfig")
		if exists("VERSION") && os.Getenv("CIRCLE_BRANCH") != "release" {
			build := fmt.Sprintf("build-%s", os.Getenv("CIRCLE_BUILD_NUM"))
			sh("echo", q(build), "> VERSION")
		}
		info := NewProjectInfo()
		path := shell.Path("/home/circleci/.go_workspace/src", info.Repo)
		sh("rm -rf", path)
		sh("mkdir -p", filepath.Dir(path))
		sh("cd .. && mv", info.Name, path, "&&", "ln -s", path, info.Name)
	},
}
