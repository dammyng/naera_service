package config

import (
	"bills/internals/config"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DSN        string
	AmqpBroker string
	JWTKey     string
	RedisHost string
	RedisPass string
	GrpcHost string
}

func NewApConfig() AppConfig {

	if os.Getenv("Environment") != "production" && os.Getenv("Environment") != "docker" {
		loadEnv()
	}

	var appConfig AppConfig
	appConfig.DSN = getDSN(config.NewDBConfig())
	appConfig.AmqpBroker = os.Getenv("AMQP_URL")
	appConfig.JWTKey = os.Getenv("JWTKey")
	appConfig.RedisHost = os.Getenv("Redis_Host")
	appConfig.RedisPass = os.Getenv("RedisPass")
	appConfig.GrpcHost = os.Getenv("BILL_GRPC_PORT")
	return appConfig
}

func getDSN(db config.DBConfig) string {
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
	err := godotenv.Load(os.Getenv("CMD_PATH")+".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}
