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
				"id":"35635d-1fc52687bbkb","beneficiary":"08069475323","provider":"Airtime NG","amount":2000},
				{
					"id":"35635d-4b36-7b0a","beneficiary":"09069475323","provider":"Airtime NG","amount":2000}
				
			
			]`,
	
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


func TestChargeCard(t *testing.T) {
	bill := models.Bill{
		Cart:          `[
			{
				"id":"35635d-1fc52687bbkb","beneficiary":"08069475323","provider":"Airtime NG","amount":2000},
				{
					"id":"35635d-4b36-7b0a","beneficiary":"09069475323","provider":"Airtime NG","amount":2000}
			]`,
			CardId: "fwrg-hthb-thntn",
			Title: "Charge Card",
	}

	_request, err := json.Marshal(&bill)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonStr = []byte(string(_request))
	req, _ := http.NewRequest("POST", "/v1/bills/chargecard", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "bearer "+os.Getenv("test_token"))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	//log.Println(response)
	checkResponse(t, http.StatusCreated, response)
	//require.Equal()
}

func TestBillTransactions(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/bills/48793121-d6d4-42d8-b2dc-a70aa629c9b6/transactions", nil)
	req.Header.Set("Authorization", "bearer "+os.Getenv("test_token"))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	//log.Println(response)
	checkResponse(t, http.StatusOK, response)
	//require.Equal()
}

func TestBillTransactionOrders(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/bills/48793121-d6d4-42d8-b2dc-a70aa629c9b6/transaction/ac4de857-d43c-4f63-a68c-5d04658b9aef", nil)
	req.Header.Set("Authorization", "bearer "+os.Getenv("test_token"))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	//log.Println(response)
	checkResponse(t, http.StatusOK, response)
	//require.Equal()
}

func TestTransactionOrders(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/biller/transactions", nil)
	req.Header.Set("Authorization", "bearer "+os.Getenv("test_token"))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	//log.Println(response)
	checkResponse(t, http.StatusOK, response)
	//require.Equal()
}