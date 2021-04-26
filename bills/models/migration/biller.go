package migration

type Biller struct {
	Id        string `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	CardToken string `gorm:"null;unique" json:"card_token,omitempty"`
	Cart      string `json:"cart,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	DeletedAt int64  `json:"deleted_at,omitempty"`
}