package migration

type Transaction struct {
	Id            string  `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	FlRef         string  `json:"flRef,omitempty"`
	Biller        string  `gorm:"not null" json:"biller,omitempty"`
	Bill          string  `gorm:"not null" json:"bill,omitempty"`
	Title         string  `json:"title,omitempty"`
	BillingMethod string  `json:"billingMethod,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
	TransRef      string  `json:"transRef,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
}
