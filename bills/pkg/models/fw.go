package models

type FwBillCategory struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data []FwBill `json:"data"`
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

func (fwBill *FwBill) ToDefaults(provider string) DisplayBill  {
	return DisplayBill{
		Provider: provider,
		Title: fwBill.ShortName,
	}
}
