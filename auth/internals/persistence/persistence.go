package persistence

import "naerarshared/models"

type NaerarAuthDBHandler interface {
	FindUserAccount(*models.UserAccount) (*models.UserAccount, error)
	CreateUserAccount(*models.UserAccount) (string, error)
	GetUserAccount(map[string]interface{}) ([]models.UserAccount, error)
	UpdateUserAccount(*models.UserAccount, map[string]interface{}) error
}
