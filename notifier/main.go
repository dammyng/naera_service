package main

import (
	"log"
	"notifier/auth"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("APP_ENV") != "production" {
		loadEnv()
	}
	auth.StartAuthenticationListener(os.Getenv("AMQP_URL"), os.Getenv("Exchange"), os.Getenv("Queue"))
}

func loadEnv() {
	log.Println("Notifier env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
