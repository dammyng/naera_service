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

type Card struct {
	Id          string `json:"id,omitempty"`
	Token       string `json:"token,omitempty"`
	Email       string `json:"email,omitempty"`
	Status      string `json:"status,omitempty"`
	LastDigits  string `json:"lastDigits,omitempty"`
	FirstDigits string `json:"firstDigits,omitempty"`
	Provider    string `json:"provider,omitempty"`
	AddedBy     string `json:"addedBy,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
	DeletedAt   int64  `json:"deleted_at,omitempty"`
	Expires     string `json:"expires,omitempty"`
}
