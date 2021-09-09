package persistence

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("Environment") != "production" {
		LoadEnv()
	}
}

type DBConfig struct {
	Hosts    string
	Database string
	Username string
	Password string
	Port     string
}

func Config(db *DBConfig) string {

	db.Hosts = os.Getenv("DBHost")
	db.Password = os.Getenv("DBPassword")
	db.Port = os.Getenv("DBPort")
	db.Username = os.Getenv("DBUser")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC",
		db.Hosts,
		db.Username,
		db.Password,
		db.Database,
		db.Port,
	)
	return dsn
}

func InitAuthDbConfig() DBConfig {
	var dBConfig DBConfig
	dBConfig.Database = os.Getenv("NaerarAuthDb")
	return dBConfig
}

func LoadEnv() {
	log.Println("env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}
