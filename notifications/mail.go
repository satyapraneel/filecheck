package notifications

import (
	"bytes"
	"html/template"
	"net/smtp"
	"strings"

	"filecheck/config"
)

//Request struct
type MailSettings struct {
	to       []string
	subject  string
	body     string
	mailType string
}

func NewMail(to []string, subject, body, mailType string) *MailSettings {
	return &MailSettings{
		to:       to,
		subject:  subject,
		body:     body,
		mailType: mailType,
	}
}

func (r *MailSettings) SendEmail() (bool, error) {
	mailConfig := config.GetMailConfig()
	mailType := r.mailType
	if mailType == "" {
		mailType = "text/html"
	}
	auth := smtp.PlainAuth("", mailConfig.User, mailConfig.Password, mailConfig.Host)
	mime := "MIME-version: 1.0;\nContent-Type: " + mailType + "; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "\r\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := mailConfig.Host + ":" + mailConfig.Port

	if err := smtp.SendMail(addr, auth, mailConfig.From, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *MailSettings) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	r.body = strings.Replace(r.body, "\n\n\n", "<br>", -1)
	return nil
}
