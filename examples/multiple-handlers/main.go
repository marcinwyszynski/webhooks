package main

import (
	"fmt"
	"strconv"

	"github.com/joeybloggs/webhooks"
	"github.com/joeybloggs/webhooks/github"
)

const (
	path = "/webhooks"
	port = 80
)

func main() {
	hook := github.New(&github.Config{Secret: "MyGitHubSuperSecretSecrect...?"})
	hook.RegisterEvents(HandleRelease, github.ReleaseEvent)
	hook.RegisterEvents(HandlePullRequest, github.PullRequestEvent)

	err := webhooks.Run(hook, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleRelease handles GitHub release events
func HandleRelease(payload interface{}) {

	fmt.Println("Handling Release")

	pl := payload.(github.ReleasePayload)

	// only want to compile on full releases
	if pl.Release.Draft || pl.Release.Prelelease || pl.Release.TargetCommitish != "master" {
		return
	}

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

// HandlePullRequest handles GitHub pull_request events
func HandlePullRequest(payload interface{}) {

	fmt.Println("Handling Pull Request")

	pl := payload.(github.PullRequestPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}