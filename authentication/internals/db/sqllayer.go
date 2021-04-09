package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type SqlLayer struct {
	Session *gorm.DB
}

func NewSqlLayer(dsn string) *SqlLayer {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	return &SqlLayer{Session: db}
}

func (sql *SqlLayer) CreateUser() (string, error) {
	return "nil", nil
}
