package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"githubn.com/Germanicus1/lenslocked/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err = es.ForgotPassword("peter@kerschbaumer.es", "https://lenslocked.com/reset-pw")
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}
