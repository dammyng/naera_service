package models

type FwBillCategory struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []FwBill `json:"data"`
}

type FwBill struct {
	ID                int     `json:"id"`
	BillerCode        string  `json:"biller_code"`
	Name              string  `json:"name"`
	DefaultCommission float64 `json:"default_commission"`
	DateAdded         string  `json:"date_added"`
	Country           string  `json:"country"`
	IsAirtime         bool    `json:"is_airtime"`
	BillerName        string  `json:"biller_name"`
	ItemCode          string  `json:"item_code"`
	ShortName         string  `json:"short_name"`
	Fee               int     `json:"fee"`
	CommissionOnFee   bool    `json:"commission_on_fee"`
	LabelName         string  `json:"label_name"`
	Amount            float64 `json:"amount"`
}

func (fwBill *FwBill) ToDefaults(provider string) DisplayBill {
	return DisplayBill{
		Provider: provider,
		Title:    fwBill.ShortName,
	}
}

type VerifiedBill struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ResponseCode    string      `json:"response_code"`
		Address         interface{} `json:"address"`
		ResponseMessage string      `json:"response_message"`
		Name            string      `json:"name"`
		BillerCode      string      `json:"biller_code"`
		Customer        string      `json:"customer"`
		ProductCode     string      `json:"product_code"`
		Email           interface{} `json:"email"`
		Fee             int         `json:"fee"`
		Maximum         int         `json:"maximum"`
		Minimum         int         `json:"minimum"`
	} `json:"data"`
}

type OrderRequest struct{
	ItemCode    string      `json:"item_code"`
	Customer        string      `json:"customer"`
	BillerCode        string      `json:"biller_code"`
}
