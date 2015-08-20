package main

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
	"github.com/progrium/go-shell"
	"github.com/spf13/cobra"
)

func init() {
	containerCmd.AddCommand(containerUpCmd)
	containerCmd.AddCommand(containerDownCmd)
	Glu.AddCommand(containerCmd)
}

var containerCmd = &cobra.Command{
	Use: "container",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var containerUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Creates a glu container",
	Run: func(cmd *cobra.Command, args []string) {
		defer shell.ErrExit()
		shell.Trace = true
		shell.Tee = os.Stdout

		if dockerExistsByName("glu") {
			sh("docker rm -f glu")
		}
		var artifactsMount string
		if os.Getenv("CIRCLE_ARTIFACTS") != "" {
			artifactsMount = os.Getenv("CIRCLE_ARTIFACTS") + ":"
		}
		sh("docker run -d --name glu",
			"--label glu",
			"--entrypoint /usr/bin/tail",
			"--volume $PWD:/project",
			fmt.Sprintf("--volume %s/artifacts", artifactsMount),
			"gliderlabs/glu -f /dev/null")

	},
}

var containerDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Destroys a glu container",
	Run: func(cmd *cobra.Command, args []string) {
		defer shell.ErrExit()
		shell.Trace = true
		shell.Tee = os.Stdout
		if dockerExistsByName("glu") {
			sh("docker rm -f glu")
		}
	},
}

func dockerExistsByName(container string) bool {
	client, err := docker.NewClientFromEnv()
	fatal(err)
	containers, err := client.ListContainers(docker.ListContainersOptions{})
	fatal(err)
	for _, cntr := range containers {
		for _, name := range cntr.Names {
			if name[1:] == container {
				return true
			}
		}
	}
	return false
}
