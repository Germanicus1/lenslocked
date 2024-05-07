package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"githubn.com/Germanicus1/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "c8a92554e741aa"
	password = "8320480c60858d"
)

func main() {
	email := models.Email{
		From:      "test@lenslocked.com",
		To:        "peter@kerschbaumer.es",
		Subject:   "This is a test email",
		Plaintext: "This is just the plain text body",
		HTML:      "<h1>This is the HTML body off my email</h1>",
	}
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.Send(email)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}
