package restclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"bills/pkg/helpers"
)

func VerifyFwTransaction(txRef string) (*FlwVerifiedTransaction, error) {
	// jwt authentication
	reqURL := fmt.Sprintf("/transactions/%v/verify", txRef)

	flutterReq, err := helpers.BuildFlutterWaveRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	result, err := HttpReq(flutterReq)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var response FlwVerifiedTransaction
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status == "error" {
		return nil, flutterError(response.Message)
	}
	return &response, err
}

func ServiceTransaction(body string) (*ServicedTransaction, error) {
	reqURL := "/bills"

	flutterReq, err := helpers.BuildFlutterWaveRequest("POST", reqURL, ioutil.NopCloser(strings.NewReader(body)))
	if err != nil {
		return nil, err
	}

	result, err := HttpReq(flutterReq)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	var response ServicedTransaction
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	log.Println(response)

	if response.Status == "error" {
		return nil, flutterError(response.Message)
	}
	return &response, err
}

func ChargeCard(body string) (*FlwVerifiedTransaction, error) {
	flutterReq, err := helpers.BuildFlutterWaveRequest("POST", "/tokenized-charges", ioutil.NopCloser(strings.NewReader(body)))
	if err != nil {
		return nil, err
	}
	result, err := HttpReq(flutterReq)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response FlwVerifiedTransaction
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status == "error" || response.Status == "Application error" {
		return nil, flutterError(response.Message)
	}
	fmt.Println(response)
	return &response, err
}
