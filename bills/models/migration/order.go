package migration

type Order struct {
	Id            string  `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	Title         string  `json:"title,omitempty"`
	TransactionId string  `json:"transactionId,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
	Charged       bool    `gorm:"not null" json:"charged,omitempty"`
	Fulfilled     bool    `gorm:"not null" json:"fulfilled,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
}
