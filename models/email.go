package models

import "github.com/go-mail/mail/v2"

const (
	DefaultSender = "peter@kerschbaumer.es"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService(config SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
	return &es
}

type EmailService struct {
	DefaultSender string
	dialer        *mail.Dialer
}
