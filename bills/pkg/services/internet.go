package services

func GetAllInternet() (interface{}, error) {
	fwPath := "/bill-categories?internet=1"
	res, err := FWBillsHandler(fwPath)
	if err != nil {
		return nil, err
	}
	return res, err
}