package repository

import (
	"bytes"
	"go-app/src/domain/entity"
	"go-app/src/infrastructure/configs"
	"strconv"
	"text/template"

	"gopkg.in/gomail.v2"
)

func NewEmailSender() EmailSender {
	return &EmailSenderImpl{}
}

type EmailSender interface {
	Send(e *entity.Email)
}

type EmailSenderImpl struct{}
var contentType = "text/plain"

func (es EmailSenderImpl) Send(e *entity.Email) {
	smtpHost := configs.Config.Mail.Host
	smtpPort, _ := strconv.Atoi(configs.Config.Mail.Port)
	smtpUsername := configs.Config.Mail.User
	smtpPassword := configs.Config.Mail.Password
	t, err := template.ParseFiles(e.TemplateFiles...)
	buff := &bytes.Buffer{}
	t.Execute(buff, e.TemplateVars)
  if err != nil {
    panic(err)
  }

	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", e.Receivers...)
	m.SetHeader("Subject", e.Subject)
	m.SetBody(contentType, buff.String())

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
