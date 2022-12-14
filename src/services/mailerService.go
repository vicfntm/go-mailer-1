package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/vcfntm/go-mailer-1/src/models"
)

type MailerService struct {
	emailLogin string
	password   string
	host       string
	port       string
	adminEmail string
}

var SUBJECT = "Subject: client feedback form \n"
var MIME_TYPE = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"

func (mailer *MailerService) Push(pack models.FeedbackForm) (bool, error) {

	target := []string{mailer.adminEmail}

	body := fmt.Sprintf("name: %s \n\n phone: %s \n\n text: %s \n", pack.Name, pack.Phone, pack.MessageText)
	_, err := mailer.SendEmailSMTP(target, body)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}

func (mailer *MailerService) SendEmailSMTP(to []string, body string) (bool, error) {

	emailAuth := smtp.PlainAuth("", mailer.emailLogin, mailer.password, mailer.host)

	msg := []byte(SUBJECT + MIME_TYPE + "\n" + body)
	addr := fmt.Sprintf("%s:%s", mailer.host, mailer.port)

	if err := smtp.SendMail(addr, emailAuth, mailer.emailLogin, to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func NewMailerService() *MailerService {
	return &MailerService{
		emailLogin: os.Getenv("EMAIL_LOGIN"),
		password:   os.Getenv("EMAIL_PASSWORD"),
		host:       os.Getenv("EMAIL_HOST"),
		port:       os.Getenv("EMAIL_PORT"),
		adminEmail: os.Getenv("EMAIL_ADMIN"),
	}
}
