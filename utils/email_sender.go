package utils

import (
	"context"
	"fmt"
	"mime"

	"gopkg.in/gomail.v2"
)

type EmailSenderConfig struct {
	ServerHost  string
	ServerPort  int
	SenderEmail string
	SenderName  string
	Password    string
}

func SendEmail(ctx context.Context, conf EmailSenderConfig, receiver string, subj string, body string) error {
	m := gomail.NewMessage()
	fromName := mime.QEncoding.Encode("utf-8", conf.SenderName)
	m.SetHeader("From", fmt.Sprintf("%s <%s>", fromName, conf.SenderEmail))
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", subj)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(conf.ServerHost, conf.ServerPort, conf.SenderEmail, conf.Password)
	d.SSL = true
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
