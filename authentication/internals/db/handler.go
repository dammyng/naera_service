package db

type Handler interface {
	CreateUser() (string, error)
}
