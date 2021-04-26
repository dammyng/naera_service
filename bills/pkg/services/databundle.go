package services


func GetAllDataBundle() (interface{}, error) {
	fwPath := "/bill-categories?data_bundle=1"
	res, err := FWBillsHandler(fwPath)
	if err != nil {
		return nil, err
	}
	return res, err
}