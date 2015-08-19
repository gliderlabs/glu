package main

import (
	"os"
	"strings"

	"github.com/progrium/go-shell"
	"github.com/spf13/cobra"
)

func init() {
	Glu.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build <name> <os-list> [pkgs]",
	Short: "Builds a Go project of Glider Labs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
			os.Exit(1)
		}
		var (
			name   = args[0]
			osList = strings.Split(args[1], ",")
			pkgs   = optArg(args, 2, ".")
		)

		defer shell.ErrExit()
		shell.Trace = true
		shell.Tee = os.Stdout

		os.Setenv("CGO_ENABLED", "0")
		for i := range osList {
			os.Setenv("GOOS", strings.ToLower(osList[i]))
			path := shell.Path("build", strings.Title(osList[i]))
			sh("mkdir -p", path)
			sh("go build -a -installsuffix cgo -o", shell.Path(path, name), pkgs)
		}
	},
}
