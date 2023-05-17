package repository

import (
	"fmt"
	"go-app/src/domain/entity"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func NewEmailSender() EmailSender {
	return &EmailSenderImpl{}
}

type EmailSender interface {
	Send(e *entity.Email)
}

type EmailSenderImpl struct{}

func (es EmailSenderImpl) Send(e *entity.Email) {
	godotenv.Load()
	smtpServer := fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT"))
	fmt.Println(smtpServer)
	auth := smtp.CRAMMD5Auth(os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"))
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(e.Receivers, ","), e.Subject, e.Body))

	if err := smtp.SendMail(smtpServer, auth, e.From, e.Receivers, msg); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
