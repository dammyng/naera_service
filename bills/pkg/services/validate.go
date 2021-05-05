package services

import (
	"bills/models"
	"fmt"
)

func ValidateOrderItem(order models.OrderRequest) (*models.VerifiedBill, error) {
	rs, err := FWVerifyBillsHandler(fmt.Sprintf("/bill-items/%s/validate?code=%s&customer=%s", order.ItemCode, order.BillerCode, order.Customer))
	if err != nil {
		return nil, err
	}
	return rs, err
}
