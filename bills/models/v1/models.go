package models

type DisplayCategory struct {
	Title     string `json:"title"`
	CreatedOn string `json:"created_on"`
	IsActive  bool   `json:"is_active"`
}
