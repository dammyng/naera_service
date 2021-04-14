package db

import "authentication/models/v1"

type Handler interface {
	CreateUser(*models.Account) (string, error)
	FindUser(*models.Account) (*models.Account, error)
	UpdateUser(*models.Account, *models.Account) error
	UpdateUserMap(*models.Account, map[string]interface{}) error
}
