package services

import (
	"bills/models"
	"bills/pkg/helpers"
	"encoding/json"
)

func FWBillsHandler(path string) ([]models.FwBill, error) {
	var result []models.FwBill
	req, err := helpers.BuildFlutterWaveRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	res, err := helpers.NetClient.Do(req)
	if err != nil {
		return nil, err

	}
	var fwCart models.FwBillCategory
	err = json.NewDecoder(res.Body).Decode(&fwCart)
	if err != nil {
		return nil, err
	}
	for _, v := range fwCart.Data {
		if v.Country == "NG"{
			result = append(result, v) 
		}
	}
	return result, err
}

func FWVerifyBillsHandler(path string) (*models.VerifiedBill, error) {
	var result models.VerifiedBill
	req, err := helpers.BuildFlutterWaveRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}
	res, err := helpers.NetClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}