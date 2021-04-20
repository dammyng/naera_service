package models

var FW_BILL_PAYMENT = `{
	"country": "NG",
	"customer": "+23490803840303",
	"amount": 500,
	"recurrence": "ONCE",
	"type": "AIRTIME",
	"reference": "9300049404444"
 }`

type DisplayBill struct {
	Id                int     `json:"id"`
	Provider          string  `json:"provider"`
	Amount          string  `json:"amount"`
	Title          string  `json:"title"`
	Caption          string  `json:"caption"`


}

type FwBill struct {
	Id                int     `json:"id"`
	BillerCode        string  `json:"biller_code"`
	DefaultCommission float32 `json:"default_commission"`
	Name              string  `json:"name"`
	Country           string  `json:"country"`
	DateAdded         string  `json:"date_added"`
	IsAirtime         bool    `json:"is_airtime"`
	BillerName        string  `json:"biller_name"`
	ItemCode          string  `json:"item_code"`
	ShortName         string  `json:"short_name"`
	Fee               float32 `json:"fee"`
	CommissionOnFee   bool    `json:"commission_on_fee"`
	LabelName         string  `json:"label_name"`
	Amount            float32 `json:"amount"`
}
