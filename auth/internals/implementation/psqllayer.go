package implementation

import (
	"log"
	"naerarshared/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqlLayer struct {
	NaerarUserDb *gorm.DB
}

func NewPsqlLayer(userDsn string) *PsqlLayer {
	userDb, err := gorm.Open(postgres.Open(userDsn), &gorm.Config{})
	if err != nil {
		log.Panicf("FAILED TO OPEN DATABASE: %v", err)
	}
	return &PsqlLayer{NaerarUserDb: userDb}
}

func (sql *PsqlLayer) CreateUserAccount(newUser *models.UserAccount) (string, error) {
	session := sql.NaerarUserDb
	return newUser.Id, session.Create(&newUser).Error
}

func (sql *PsqlLayer) FindUserAccount(user *models.UserAccount) (*models.UserAccount, error) {
	session := sql.NaerarUserDb
	response := models.UserAccount{}
	err := session.Where(user).First(&response).Error
	return &response, err
}

func (sql *PsqlLayer) GetUserAccount(arg map[string]interface{}) ([]models.UserAccount, error) {
	session := sql.NaerarUserDb
	response := []models.UserAccount{}
	err := session.Where(arg).Error
	return response, err
}

func (sql *PsqlLayer) UpdateUserAccount(user *models.UserAccount, dict map[string]interface{}) error {
	session := sql.NaerarUserDb
	return session.Model(&user).Updates(dict).Error
}
