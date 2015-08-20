package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	Glu.AddCommand(infoCmd)
}

type ProjectInfo struct {
	Name    string
	Owner   string
	Repo    string
	Version string
}

func NewProjectInfo() ProjectInfo {
	repo := repoLocation()
	return ProjectInfo{
		Name:    filepath.Base(repo),
		Owner:   filepath.Base(filepath.Dir(repo)),
		Repo:    repo,
		Version: findVersion(),
	}
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show info about project",
	Run: func(cmd *cobra.Command, args []string) {
		info := NewProjectInfo()
		fmt.Println("Name:", info.Name)
		fmt.Println("Version:", info.Version)
		fmt.Println("Owner:", info.Owner)
		fmt.Println("Repo:", info.Repo)
	},
}
