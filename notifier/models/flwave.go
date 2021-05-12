package models
type ServicedTransaction struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Name        string  `json:"name"`
		Network     string  `json:"network"`
		Amount      float64 `json:"amount"`
		PhoneNumber string  `json:"phone_number"`
		TxRef       string  `json:"tx_ref"`
		FlwRef      string  `json:"flw_ref"`
	} `json:"data"`
}

type ServiceRequestPayload struct {
	Country    string  `json:"country"`
	Customer   string  `json:"customer"`
	Amount     float64 `json:"amount"`
	Recurrence string  `json:"recurrence"`
	Type       string  `json:"type"`
	Reference  string  `json:"reference"`
}