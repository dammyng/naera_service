package db

import (
	"bills/models/v1"
	"errors"
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

func (sql *SqlLayer) CreateABiller(user *models.Biller) (string, error) {
	err := sql.Session.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.Id, err
}

func (sql *SqlLayer) FindABiller(arg *models.Biller) (*models.Biller, error) {
	session := sql.Session
	var dA models.Biller
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateABiller(old *models.Biller, new *models.Biller) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}


func (sql *SqlLayer) UpdateABillerMap(arg *models.Biller, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}