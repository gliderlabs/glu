package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/progrium/go-shell"
	"github.com/spf13/cobra"
)

var (
	sh = shell.Run
	q  = shell.Quote
)

var Glu = &cobra.Command{
	Use:   "glu",
	Short: "glu is a collection of utility commands for Glider Labs projects",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	Glu.Execute()
}

func fatal(err error) {
	if err != nil {
		fmt.Println("!!", err)
		os.Exit(1)
	}
}

func fatalMsg(err error, msg string) {
	if err != nil {
		fmt.Println("!!", msg)
		os.Exit(1)
	}
}

func optArg(args []string, i int, default_ string) string {
	if i+1 > len(args) {
		return default_
	}
	return args[i]
}

func exists(path ...string) bool {
	_, err := os.Stat(filepath.Join(path...))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fatal(err)
	return true
}

func writeFile(path, data string) {
	fatal(ioutil.WriteFile(path, []byte(strings.Trim(data, "\n")+"\n"), 0644))
}

func mkdirAll(path ...string) {
	fatal(os.MkdirAll(filepath.Join(path...), 0777))
}

func shellOutput(cmd string) string {
	args := strings.Split(cmd, " ")
	out, _ := exec.Command(args[0], args[1:]...).Output()
	return strings.Trim(string(out), " \n")
}

func repoPath() string {
	if insideContainer() {
		os.Chdir("/project")
	}
	repo := shellOutput("git config --get remote.origin.url")
	repo = strings.TrimPrefix(repo, "http://")
	repo = strings.TrimPrefix(repo, "https://")
	repo = strings.TrimPrefix(repo, "git@")
	repo = strings.TrimSuffix(repo, ".git")
	return repo
}
