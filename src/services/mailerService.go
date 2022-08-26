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
}

func (mailer *MailerService) Push(pack models.FeedbackForm) (bool, error) {

	target := []string{"khrestyk.shop@gmail.com"}

	body := fmt.Sprintf("name: %s \n\n phone: %s \n\n text: %s \n", pack.Name, pack.Phone, pack.MessageText)
	res, err := mailer.SendEmailSMTP(target, body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)

	return true, nil
}

func (mailer *MailerService) SendEmailSMTP(to []string, body string) (bool, error) {

	emailAuth := smtp.PlainAuth("", mailer.emailLogin, mailer.password, mailer.host)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"
	msg := []byte(subject + mime + "\n" + body)
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
	}
}
