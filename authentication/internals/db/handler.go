package db

import "authentication/models/v1"

type Handler interface {
	CreateUser(*models.Account) (string, error)
}
