package migration

type Bill struct {
	Id              string `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	Biller          string `gorm:"not null" json:"biller,omitempty"`
	By              string `gorm:"not null" json:"by,omitempty"`
	Cart            string `json:"cart,omitempty"`
	Title           string `json:"title,omitempty"`
	TransactionId   string `json:"transaction_id,omitempty"`
	Reoccurring     bool   `json:"reoccurring,omitempty"`
	NextPaymentDate int64  `json:"nextPaymentDate,omitempty"`
	Active          bool   `json:"active,omitempty"`
	PayingWith      int    `json:"payingWith,omitempty"`
	CreatedAt       int64  `json:"created_at,omitempty"`
	UpdatedAt       int64  `json:"updated_at,omitempty"`
	DeletedAt       int64  `json:"deleted_at,omitempty"`
	LastPaymentDate       int64  `json:"last_payment_date,omitempty"`
}
