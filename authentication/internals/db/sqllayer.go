package db

import (
	"authentication/models/v1"
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

func (sql *SqlLayer) CreateUser(user *models.Account) (string, error) {
	err := sql.Session.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.Id, err
}

func (sql *SqlLayer) FindUser(arg *models.Account) (*models.Account, error) {
	session := sql.Session
	var dA models.Account
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) FindUsers(arg string) ([]*models.Account, error) {
	session := sql.Session
	var dA []*models.Account
	log.Println(arg)
	err := session.Where("user_name LIKE ?", arg + "%").Find(&dA).Error
	if err != nil {
		return nil, err
	}
	log.Println(dA)
	return dA, err
}

func (sql *SqlLayer) UpdateUser(old *models.Account, new *models.Account) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}


func (sql *SqlLayer) UpdateUserMap(arg *models.Account, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}
