package config

import (
	"authentication/internals/db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DSN        string
	AmpqBroker string
	JWTKey     string
	ReddisHost string
	ReddisPass string
}

func NewApConfig() AppConfig {

	if os.Getenv("Environment") != "production" {
		loadEnv()
	}

	var appConfig AppConfig
	appConfig.DSN = getDSN(db.NewDBConfig())
	appConfig.AmpqBroker = os.Getenv("AmpqHost")
	appConfig.JWTKey = os.Getenv("JWTKey")
	appConfig.ReddisHost = os.Getenv("ReddisHost")
	appConfig.ReddisPass = os.Getenv("ReddisPass")
	return appConfig
}

func getDSN(db db.DBConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Hosts,
		db.Port,
		db.Database)
	return dsn
}

func loadEnv() {
	log.Println("env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}
