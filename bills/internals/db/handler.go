package db

import "bills/models/v1"


type Handler interface {
	CreateABiller(*models.Biller) (string, error)
	FindABiller(*models.Biller) (*models.Biller, error)
	UpdateABiller(*models.Biller, *models.Biller) error
	UpdateABillerMap(*models.Biller, map[string]interface{}) error
	//AddCategory(*models.DisplayCategory) (string, error)
	//UpdateCategory(*models.DisplayCategory, *models.DisplayCategory) error
	//UpdateCategoryMap(*models.DisplayCategory, map[string]interface{}) error
}
