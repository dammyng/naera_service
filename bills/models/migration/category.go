
package migration


type BillCategory struct {
	Id              string  `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	RefTitle        string  `gorm:"not null;unique" json:"refTitle,omitempty"`
	DisplayTitle    string  ` json:"displayTitle,omitempty"`
	CurrentDiscount float32 `json:"currentDiscount,omitempty"`
	Active          bool    `json:"active,omitempty"`
	CreatedAt       int64   `json:"created_at,omitempty"`
	UpdatedAt       int64   `json:"updated_at,omitempty"`
}