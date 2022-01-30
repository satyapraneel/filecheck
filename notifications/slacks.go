package notifications

import (
	"filecheck/config"
	"fmt"
	"log"
	"strings"

	"github.com/slack-go/slack"
)

func SendSlackMessage(message string, errorContext map[string]string) {
	errorContextString := formatErrorContext(errorContext)
	p := config.GetSlackConfig()
	if errorContext["code"] != "" {
		message = errorContext["code"] + " : " + message
	}

	attachment := slack.Attachment{
		Title:      "Error",
		Fallback:   message,
		Text:       message,
		Color:      "#FF2323",
		MarkdownIn: []string{"fields"},
		Footer:     "Powered By GO",
		FooterIcon: "https://emojis.slackmojis.com/emojis/images/1454546974/291/golang.png",
		Fields: []slack.AttachmentField{
			{
				Title: "Context",
				Value: errorContextString,
				Short: false,
			},
		},
	}

	option := slack.MsgOption(
		slack.MsgOptionAttachments(attachment),
	)

	userName := slack.MsgOptionUsername("DA Slack")

	_, _, _, err := p.SlackClient.SendMessage(p.Channel, option, userName)

	if err != nil {
		log.Printf("ERROR Connecting to slack : %v", err)
	}
}

func formatErrorContext(errorContext map[string]string) string {
	errorContextString := ""

	for key, value := range errorContext {
		errorContextString += fmt.Sprintf("\n%s: \n%s", key, strings.Replace(value, "\n", "\t", -1))
	}
	return errorContextString
}
