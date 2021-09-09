package models

type DisplayBill struct {
	Id        int    `json:"id"`
	Provider  string `json:"provider"`
	Amount    string `json:"amount"`
	Title     string `json:"title"`
	ShortName string `json:"short_name"`
}

type DisplayVerifyBill struct {
	Id          string    `json:"id"`
	Name        string `json:"name"`
	Amount      float64 `json:"amount"`
	Title       string `json:"title"`
	Beneficiary string `json:"beneficiary"`
	Status      string `json:"status"`
}
