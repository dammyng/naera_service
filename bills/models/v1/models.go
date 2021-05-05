package models

type DisplayCategory struct {
	Title     string `json:"title"`
	CreatedOn string `json:"created_on"`
	IsActive  bool   `json:"is_active"`
}

type ServiceRequestPayload struct {
	Country    string `json:"country"`
	Customer   string `json:"customer"`
	Amount     float64    `json:"amount"`
	Recurrence string `json:"recurrence"`
	Type       string `json:"type"`
	Reference  string `json:"reference"`
}

type InCartItem struct {
	ID          string `json:"id"`
	Beneficiary string `json:"beneficiary"`
	Provider    string `json:"provider"`
	Amount      float64 `json:"amount"`
}