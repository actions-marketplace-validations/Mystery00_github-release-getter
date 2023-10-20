package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	githubToken = os.Getenv("GITHUB_TOKEN")
	repoOwner   = os.Args[1]
	repoName    = os.Args[2]
)

func main() {
	uri := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	release := GithubRelease{}
	err = json.Unmarshal(bytes, &release)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`::set-output name=tagName::%s`, release.TagName)
}

type GithubRelease struct {
	TagName string `json:"tag_name"`
}
