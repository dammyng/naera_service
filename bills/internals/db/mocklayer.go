package db

/****
import "bills/internals/models/v1"

type MockLayer struct {
}

func NewMockLayer() *MockLayer {
	return &MockLayer{}
}

func (ml *MockLayer) GetAllCategories() ([]models.DisplayCategory, error) {
	return categories, nil
}

func (ml *MockLayer) GetLiveCategories() ([]models.DisplayCategory, error) {
	var result []models.DisplayCategory
	for _, v := range categories {
		if v.IsActive {
			result = append(result, v)
		}
	}
	return result, nil
}
**/