package main

import (
	"os"

	"github.com/catoaune/multichannel/channel/slack"
)

func main() {
	slackChannel := os.Getenv("slack_url")
	slackConfig := slack.NewConfig(slackChannel)
	slackConfig.SendNotification("Hei på deg!")
}
