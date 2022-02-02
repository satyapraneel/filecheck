package notifications

import (
	"filecheck/config"
	"fmt"
	"log"
	"strings"

	"github.com/slack-go/slack"
)

func SendSlackSuccessMessage(message string) {
	sendSlackMessage("Success", message, "", "#008000")
}

func SendSlackErrorMessage(message string, errorContext map[string]string) {
	errorContextString := formatContext(errorContext)
	if errorContext["code"] != "" {
		message = errorContext["code"] + " : " + message
	}
	sendSlackMessage("Error", message, errorContextString, "#FF2323")
}

func formatContext(context map[string]string) string {
	contextString := ""

	for key, value := range context {
		contextString += fmt.Sprintf("\n%s: \n%s", key, strings.Replace(value, "\n", "\t", -1))
	}
	return contextString
}

func sendSlackMessage(title string, message string, context string, colorCode string) {

	p := config.GetSlackConfig()
	attachment := slack.Attachment{
		Title:      title,
		Fallback:   message,
		Text:       message,
		Color:      colorCode,
		MarkdownIn: []string{"fields"},
		Footer:     "Powered By GO",
		FooterIcon: "https://emojis.slackmojis.com/emojis/images/1454546974/291/golang.png",
		Fields: []slack.AttachmentField{
			{
				Title: "Context",
				Value: context,
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
