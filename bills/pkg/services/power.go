package services

func GetAllPower() (interface{}, error) {
	fwPath := "/bill-categories?power=1"
	res, err := FWBillsHandler(fwPath)
	if err != nil {
		return nil, err
	}
	return res, err
}