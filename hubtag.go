package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	GluCmd.AddCommand(hubtagCmd)
}

var hubtagCmd = &cobra.Command{
	Use:   "hubtag <dockerhub-repo> <tag> [git-tag] [dockerfile-path]",
	Short: "Creates an automated build tag for a Docker Hub repo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
			os.Exit(1)
		}
		dockerhub_repo := args[0]
		tag := make(map[string]string)
		tag["name"] = args[1]
		tag["source_type"] = "Tag"
		tag["source_name"] = args[1]
		if len(args) >= 3 {
			tag["source_name"] = args[2]
		}
		tag["dockerfile_location"] = "/"
		if len(args) >= 4 {
			tag["dockerfile_location"] = args[3]
		}
		tagJson, err := json.Marshal(&tag)
		fatal(err)
		login, err := base64.StdEncoding.DecodeString(os.Getenv("DOCKERHUB_LOGIN"))
		fatalMsg(err, "Bad login value for DOCKERHUB_LOGIN")

		resp, err := http.Post("https://hub.docker.com/v2/users/login",
			"application/json", bytes.NewBuffer(login))
		fatal(err)
		defer resp.Body.Close()
		dec := json.NewDecoder(resp.Body)
		var token map[string]string
		fatal(dec.Decode(&token))

		client := new(http.Client)
		req, err := http.NewRequest("POST",
			fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/autobuild/tags/", dockerhub_repo),
			bytes.NewBuffer(tagJson))
		fatal(err)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("JWT %s", token["token"]))
		resp, err = client.Do(req)
		fatal(err)
		if resp.StatusCode != 201 {
			fmt.Println(resp)
			os.Exit(1)
		}
	},
}
