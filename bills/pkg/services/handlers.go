package services

import (
	"bills/pkg/helpers"
	"bills/pkg/models"
	"encoding/json"
)

func FWBillsHandler(path string) ([]models.DisplayBill, error) {
	var result []models.DisplayBill
	req, err := helpers.BuildFlutterWaveRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	res, err := helpers.NetClient.Do(req)
	var fwCart models.FwBillCategory
	err = json.NewDecoder(res.Body).Decode(&fwCart)
	if err != nil {
		return nil, err

	}
	for _, v := range fwCart.Data {
		if v.Country == "NG"{
			result = append(result, v.ToDefaults("fw")) 
		}
	}
	return result, err
}