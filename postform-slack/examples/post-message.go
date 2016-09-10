package main

import (
	"github.com/drgarcia1986/slack"
)

func main() {
	slackClient := slack.New("slack-token")
	err := slackClient.PostMessage("#general", "Golang", ":smile:", "Hello World")
	if err != nil {
		panic(err)
	}
}
