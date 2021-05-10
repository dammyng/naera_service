
package events

type ServiceAirTimeEvent struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
	Amount float64 `json:"amount"`
	Transaction string  `json:"transaction"`
	OrderURL string `json:"orderURL"`
}

func (e *ServiceAirTimeEvent) EventName() string {
	return "buy.airtime"
}