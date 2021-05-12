package models
import(
	"shared/amqp/events"
	"encoding/hex"
	"github.com/twinj/uuid"
	"os"

)
type DisplayCategory struct {
	Title     string `json:"title"`
	CreatedOn string `json:"created_on"`
	IsActive  bool   `json:"is_active"`
}

type ServiceRequestPayload struct {
	Country    string  `json:"country"`
	Customer   string  `json:"customer"`
	Amount     float64 `json:"amount"`
	Recurrence string  `json:"recurrence"`
	Type       string  `json:"type"`
	Reference  string  `json:"reference"`
}

type InCartItem struct {
	ID          string  `json:"id"`
	Beneficiary string  `json:"beneficiary"`
	Provider    string  `json:"provider"`
	Amount      float64 `json:"amount"`
	Category      string `json:"category"`
	Transaction      string `json:"transaction"`
}

func (i *InCartItem)CreateMsg() events.Event {

	switch i.Category {
	case "airtime":
		msg := events.ServiceAirTimeEvent{
			ID:    hex.EncodeToString(uuid.NewV4().Bytes()),
			Phone: i.Beneficiary,
			Amount: i.Amount,
			Transaction: i.Transaction,
			OrderURL: os.Getenv("CreateOrderUrl"),// "http://localhost:7777/v1/bills/createorder",
		}
		return &msg
	case "cabletv":
		return &events.ServiceAirTimeEvent{}
	default:
		return &events.ServiceAirTimeEvent{}
	}
}
