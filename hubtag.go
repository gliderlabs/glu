package main

import (
	"fmt"
	"os"

	"github.com/progrium/go-shell"
	"github.com/spf13/cobra"
)

func init() {
	Glu.AddCommand(hubtagCmd)
}

var hubtagCmd = &cobra.Command{
	Use:   "hubtag <dockerhub-repo> <tag> [git-tag] [dockerfile-path]",
	Short: "Creates an automated build tag for a Docker Hub repo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
			os.Exit(1)
		}
		dockerhubRepo := args[0]
		tagName := args[1]
		tagSource := args[1]
		if len(args) >= 3 {
			tagSource = args[2]
		}
		dockerfileLocation := "/"
		if len(args) >= 4 {
			dockerfileLocation = args[3]
		}

		defer shell.ErrExit()
		shell.Trace = true
		shell.Tee = os.Stdout
		sh("go get -u github.com/progrium/dockerhub-tag")

		sh(fmt.Sprintf("dockerhub-tag set %s %s %s %s", dockerhubRepo, tagName, tagSource, dockerfileLocation))

	},
}
