package config

import (
	"github.com/slack-go/slack"
)

type SlackConfig struct {
	Channel     string
	SlackClient *slack.Client
}

func GetSlackConfig() *SlackConfig {
	slackClient := slack.New(getEnv("SLACK_TOKEN", ""))
	return &SlackConfig{
		Channel:     getEnv("SLACK_CHANNEL", ""),
		SlackClient: slackClient,
	}
}
