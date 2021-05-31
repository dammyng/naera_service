package restclient

import "time"

type FlwVerifiedTransaction struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID                int       `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FewRef            string    `json:"few_ref"`
		Amount            float64   `json:"amount"`
		Currency          string    `json:"currency"`
		ChargedAmount     float64   `json:"charged_amount"`
		AppFee            float64   `json:"app_fee"`
		MerchantFee       float64   `json:"merchant_fee"`
		AmountSettled     float64   `json:"amount_settled"`
		ProcessorResponse string    `json:"processor_response"`
		AuthModel         string    `json:"auth_model"`
		IP                string    `json:"ip"`
		Narration         string    `json:"narration"`
		Status            string    `json:"status"`
		PaymentType       string    `json:"payment_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountID         int       `json:"account_id"`
		Meta              struct {
			Checkoutinitaddress string `json:"__CheckoutInitAddress"`
			ConsumerMac         string `json:"consumer_mac"`
		} `json:"meta"`
		Card struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Token        string `json:"token"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
		Customer struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			PhoneNumber string    `json:"phone_number"`
			Email       string    `json:"email"`
			CreatedAt   time.Time `json:"created_at"`
		} `json:"customer"`
	} `json:"data"`
}

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

type FundWalletPayload struct {
	TransactionID string  `json:"transactionID"`
	WalletID string  `json:"walletID"`
}
