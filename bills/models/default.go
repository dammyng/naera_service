package models


type DisplayBill struct {
	Id       int    `json:"id"`
	Provider string `json:"provider"`
	Amount   string `json:"amount"`
	Title    string `json:"title"`
	ShortName  string `json:"short_name"`
}
