package migration


type Transaction struct {
	Id            string  `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	Biller        string  `gorm:"not null" json:"biller,omitempty"`
	Title         string  `json:"title,omitempty"`
	BillingMethod string  `json:"billingMethod,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
	TransRef      string  `json:"transRef,omitempty"`
	Bill          string  `gorm:"not null" json:"bill,omitempty"`
	Charged          bool  `gorm:"not null" json:"charged,omitempty"`
	Served          bool  `gorm:"not null" json:"served,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
}
