package mailer

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	UserMail     string
	UserMailPass string
	Host         string
	Port         int
}

func NewMailer(user, password, host string, port int) *Mailer {
	return &Mailer{
		UserMail:     user,
		UserMailPass: password,
		Host:         host,
		Port:         port,
	}
}

func (t *Mailer) SendFiledMail(filename, to, subject, bodyType, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("To", to)
	message.SetHeader("From", t.UserMail)
	message.SetHeader("Subject", subject)
	message.Embed(filename)
	message.SetBody(bodyType, body)
	d := gomail.NewDialer(t.Host, t.Port, t.UserMail, t.UserMailPass)
	err := d.DialAndSend(message)
	fmt.Println(err, "lelkek", t.Host)
	return err
}
