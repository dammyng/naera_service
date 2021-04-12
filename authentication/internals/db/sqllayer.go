package db

import (
	"authentication/models/v1"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (sql *SqlLayer) CreateUser(user *models.Account) (string, error) {
	sql.Session.AutoMigrate(&models.Account{})
	err := sql.Session.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.Id, err
}
