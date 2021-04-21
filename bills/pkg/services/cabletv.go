
package services

func GetAllCableTv() (interface{}, error) {
	fwPath := "/bill-categories?cables=1"
	res, err := FWBillsHandler(fwPath)
	if err != nil {
		return nil, err
	}
	return res, err
}