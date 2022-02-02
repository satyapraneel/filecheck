package services

import (
	"filecheck/notifications"
	"fmt"
	"log"
)

var erroMessage = "DA : Error in file generation"
var successMessage = "DA: All file generation is success"

func (app App) SendNotification(validatedFiles ValidateFileStruct) bool {

	//send notification
	channels := app.DAConfig.NotificationType
	for _, channel := range channels {
		switch channel {
		case "email":
			println("sending email")
			if validatedFiles.SendErrorNotification {
				app.SendEmailNotification(validatedFiles, erroMessage, "file_errors.html")
			} else {
				app.SendEmailNotification(validatedFiles, successMessage, "file_success.html")
			}
		case "slack":
			if validatedFiles.SendErrorNotification {
				println("Sending Error Slack notification")
				app.SendSlackNotification(validatedFiles)
			} else {
				println("Sending Success Slack notification")
				notifications.SendSlackSuccessMessage(successMessage)
			}
		default:
			println("no channels available to send the notifications")
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
	notifications.SendSlackErrorMessage(erroMessage, errorContext)
}

func (app App) SendEmailNotification(validatedFiles ValidateFileStruct, message string, template string) {
	mail := notifications.NewMail(app.DAConfig.EmailTo, message, "", "")
	// mailTemplate := "/ui/html/mails/file_errors.html"
	mailTemplate := "/ui/html/mails/" + template
	errs := mail.ParseTemplate(mailTemplate, validatedFiles)
	if errs != nil {
		log.Printf("template parse : %v", errs)
	}
	ok, err := mail.SendEmail()
	fmt.Println(err)
	fmt.Println(ok)
}
