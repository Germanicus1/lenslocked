package main

import (
	"fmt"

	"github.com/go-mail/mail/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Using mailtrap.io
const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "c8a92554e741aa"
	password = "8320480c60858d"
)

func main() {
	from := "test@lenslocked.com"
	to := "peter@kerschbaumer.es"
	subject := "This is a test email"
	html := "<h1>This is the body off my email</h1>"
	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", "This is the body off my email")
	msg.AddAlternative("text/html", html)
	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}
