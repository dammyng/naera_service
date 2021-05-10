
package events

type ServiceAirTimeEvent struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
	Amount float64 `json:"amount"`
}

func (e *ServiceAirTimeEvent) EventName() string {
	return "buy.airtime"
}