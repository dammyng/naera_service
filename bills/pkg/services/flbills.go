package services

import "bills/models"



func GetAllFLBills() ([]models.FwBill, error) {
	fwPath := "/bill-categories"
	res, err := FWBillsHandler(fwPath)
	if err != nil {
		return nil, err
	}
	return res, err
}
