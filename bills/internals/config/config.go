package config

import "os"

type DBConfig struct {
	Hosts    string
	Database string
	Username string
	Password string
	Port     string
}

func NewDBConfig() DBConfig {
	var dBConfig DBConfig
	dBConfig.Database = os.Getenv("BILLS_DBDatabase")
	dBConfig.Hosts = os.Getenv("DBHost")
	dBConfig.Password = os.Getenv("DBPassword")
	dBConfig.Port = os.Getenv("DBPort")
	dBConfig.Username = os.Getenv("DBUser")
	return dBConfig
}