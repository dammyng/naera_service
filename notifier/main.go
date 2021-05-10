package main

import (
	"log"
	"notifier/auth"
	"notifier/service"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("Environment") != "production" {
		loadEnv()
	}
	go func() {
		service.StartServiceProcessListener(os.Getenv("AMQP_URL"), os.Getenv("Exchange"), os.Getenv("Queue"))
	}()
	auth.StartAuthenticationListener(os.Getenv("AMQP_URL"), os.Getenv("Exchange"), os.Getenv("Queue"))
}

func loadEnv() {
	log.Println("Notifier env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
