package rest

import (
	"bills/pkg/helpers"
	"bills/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)


func (handler *BillHandler) AllBills(w http.ResponseWriter, r *http.Request) {
	path := "/bill-categories"
	req, err := helpers.BuildFlutterWaveRequest("GET", path, nil)
	if err != nil {
		respondWithError(w, http.StatusOK, InternalServerError)
	}
	res, err :=  NetClient.Do(req)

	var _data map[string][]models.FwBill

	err = json.NewDecoder(res.Body).Decode(&_data)
	fmt.Println(_data)

	//field, ok := _data["data"].([]models.Bill)
	//if  !ok {
	//	http.Error(w, err.Error(), 400)
	//	return
	//}
	//fmt.Println(field)

	//respondWithJSON(w, http.StatusOK, field)

}