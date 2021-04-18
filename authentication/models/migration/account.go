package migration

type Account struct {
	Id              string `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	FirstName       string `gorm:"not null" json:"first_name,omitempty"`
	Surname         string `gorm:"not null" json:"surname,omitempty"`
	UserName        string `json:"user_name;unique,omitempty"`
	Email           string `gorm:"not null;unique" json:"email,omitempty"`
	PhoneNumber     string `gorm:"not null;unique" json:"phone_number,omitempty"`
	Dob             int64  `json:"dob,omitempty"`
	EmailVerifiedAt int64  `json:"emailVerifiedAt,omitempty"`
	PhoneVerifiedAt int64  `json:"phoneVerifiedAt,omitempty"`
	RefCode         string `json:"ref_code;unique,omitempty"`
	PinUpdatedAt    int64  `json:"pinUpdatedAt,omitempty"`
	Photo           string `json:"photo,omitempty"`
	Bvn             string `json:"bvn;unique,omitempty"`
	Password        []byte `gorm:"not null" json:"password,omitempty"`
	Pin             []byte `json:"pin,omitempty"`
	CreatedAt       int64  `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt       int64  `json:"updated_at,omitempty"`
	DeletedAt       int64  `json:"deleted_at,omitempty"`
	IsReady         bool   `gorm:"not null" json:"isReady,omitempty"`
	BvnVerifiedAt   int64  `json:"bvnVerifiedAt,omitempty"`
	NubanVerifiedAt int64  `json:"nubanVerifiedAt,omitempty"`
	BankCode        string `json:"bankCode,omitempty"`
	Nuban           string `json:"nuban;unique,omitempty"`
	Address         string `json:"address,omitempty"`
	State           string `json:"state,omitempty"`
	City            string `json:"city,omitempty"`
	IdCard          string `json:"idCard,omitempty"`
	Document        string `json:"document,omitempty"`
	UtilityBill     string `json:"utilityBill,omitempty"`
	Nin             string `json:"nin;unique,omitempty"`
}


type CleanAccount struct {
	Id              string `gorm:"primary_key;not null;unique" json:"id,omitempty"`
	FirstName       string `gorm:"not null" json:"first_name,omitempty"`
	Surname         string `gorm:"not null" json:"surname,omitempty"`
	UserName        string `json:"user_name;unique,omitempty"`
	Email           string `gorm:"not null;unique" json:"email,omitempty"`
	PhoneNumber     string `gorm:"not null;unique" json:"phone_number,omitempty"`
	Dob             int64  `json:"dob,omitempty"`
	EmailVerifiedAt int64  `json:"emailVerifiedAt,omitempty"`
	PhoneVerifiedAt int64  `json:"phoneVerifiedAt,omitempty"`
	RefCode         string `json:"ref_code;unique,omitempty"`
	Photo           string `json:"photo,omitempty"`
	Bvn             string `json:"bvn;unique,omitempty"`
	PinUpdatedAt    int64  `json:"pinUpdatedAt,omitempty"`
	CreatedAt       int64  `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt       int64  `json:"updated_at,omitempty"`
	DeletedAt       int64  `json:"deleted_at,omitempty"`
	IsReady         bool   `gorm:"not null" json:"isReady,omitempty"`
	BvnVerifiedAt   int64  `json:"bvnVerifiedAt,omitempty"`
	NubanVerifiedAt int64  `json:"nubanVerifiedAt,omitempty"`
	BankCode        string `json:"bankCode,omitempty"`
	Nuban           string `json:"nuban;unique,omitempty"`
	Address         string `json:"address,omitempty"`
	State           string `json:"state,omitempty"`
	City            string `json:"city,omitempty"`
	IdCard          string `json:"idCard,omitempty"`
	Document        string `json:"document,omitempty"`
	UtilityBill     string `json:"utilityBill,omitempty"`
	Nin             string `json:"nin;unique,omitempty"`
}
