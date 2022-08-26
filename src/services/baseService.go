package services

import "github.com/vcfntm/go-mailer-1/src/models"

type Mailer interface {
	Push(data models.FeedbackForm) (bool, error)
	SendEmailSMTP(to []string, body string) (bool, error)
}

type Service struct {
	Mailer
}

func NewService() *Service {
	return &Service{
		Mailer: NewMailerService(),
	}
}
