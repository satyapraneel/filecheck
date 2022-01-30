package services

import (
	"filecheck/notifications"
	"fmt"
	"log"
)

var erroMessage = "DA : Error in file generation"

func (app App) SendNotification(validatedFiles ValidateFileStruct) bool {
	if validatedFiles.SendNotification {
		//send notification
		channels := app.DAStruct.NotificationType
		for _, channel := range channels {
			switch channel {
			case "email":
				println("sending email")
				app.SendEmailNotification(validatedFiles)
			case "slack":
				println("sending Slack notification")
				app.SendSlackNotification(validatedFiles)
			default:
				println("no channels available to send the notifications")
			}
		}
	}
	return true
}

func (app App) SendSlackNotification(validatedFiles ValidateFileStruct) {
	errorContext := make(map[string]string)
	if validatedFiles.FileNotFound != "" {
		errorContext["file_not_found"] = validatedFiles.FileNotFound
	}
	if validatedFiles.FileSizeLess != "" {
		errorContext["file_threashold_less"] = validatedFiles.FileSizeLess
	}
	if validatedFiles.FileInvalidInterval != "" {
		errorContext["invalid_interval"] = validatedFiles.FileInvalidInterval
	}
	notifications.SendSlackMessage(erroMessage, errorContext)
}

func (app App) SendEmailNotification(validatedFiles ValidateFileStruct) {
	mail := notifications.NewMail(app.DAStruct.EmailTo, erroMessage, "", "")
	mailTemplate := "/ui/html/mails/file_errors.html"
	errs := mail.ParseTemplate(mailTemplate, validatedFiles)
	if errs != nil {
		log.Printf("template parse : %v", errs)
	}
	ok, err := mail.SendEmail()
	fmt.Println(err)
	fmt.Println(ok)
}
