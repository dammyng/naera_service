package db

import "bills/internals/models/v1"

type Handler interface {
	GetLiveCategories() ([]models.DisplayCategory, error)
	GetAllCategories() ([]models.DisplayCategory, error)
	//AddCategory(*models.DisplayCategory) (string, error)
	//UpdateCategory(*models.DisplayCategory, *models.DisplayCategory) error
	//UpdateCategoryMap(*models.DisplayCategory, map[string]interface{}) error
}
