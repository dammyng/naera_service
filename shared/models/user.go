package models

type UserAccount struct {
	Id        string `json:"id"  gorm:"primary_key;index"`
	Firstname string `json:"firstname" gorm:"size:255;"`
	Lastname  string `json:"lastname" gorm:"size:255;"`
	Email     string `json:"email" gorm:"primary_key;index"`
	Password  []byte `json:"-"`

	CreatedAt int64 `json:"created_at" `
	UpdatedAt int64 `json:"updated_at" `
	DeletedAt int64 `json:"deleted_at" `
}
