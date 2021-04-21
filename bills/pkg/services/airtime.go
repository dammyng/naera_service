package services



func GetAllAirtime() (interface{}, error) {
	fwPath := "/bill-categories?airtime=1"
	res, err := FWBillsHandler(fwPath)
	if err != nil {
		return nil, err
	}
	return res, err
}

func PossibleCombinations()  {
	
}
