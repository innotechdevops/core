package email

import (
	m "github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

type Utility interface {
	Send(mail *Mail) error
}

type utility struct {
}

func (u *utility) Send(mail *Mail) error {
	e := m.Email{
		To:      mail.To,
		From:    mail.From,
		Subject: mail.Subject,
		HTML:    mail.HTML,
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(
		"smtp.gmail.com:587",
		smtp.PlainAuth(
			"",
			mail.From, mail.FromPassword,
			"smtp.gmail.com",
		),
	)
	return err
}

func New() Utility {
	return &utility{}
}
