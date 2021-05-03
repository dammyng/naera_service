package rest_test

import (
	"bills/models/v1"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	//"github.com/stretchr/testify/require"
)

func TestCreateBill(t *testing.T) {
	bill := models.Bill{
		Cart:          `[
			{
				"id":"35635d-1fc52687bbkb","beneficiary":"08069475323","provider":"Airtime NG","amount":10},
				{
					"id":"35635d-4b36-7b0a","beneficiary":"09069475323","provider":"Airtime NG","amount":10}
				
			
			]`,
		TransactionId: `420012877`,
	}

	_request, err := json.Marshal(&bill)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonStr = []byte(string(_request))
	req, _ := http.NewRequest("POST", "/v1/bills/createbill", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "bearer "+os.Getenv("test_token"))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	//log.Println(response)
	checkResponse(t, http.StatusCreated, response)
	//require.Equal()
}
