package restclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func VerifyFwTransaction(key, txRef string) (*FlwVerifiedTransaction, error) {
	// jwt authentication
	reqURL, _ := url.Parse(fmt.Sprintf("https://api.flutterwave.com/v3/transactions/%v/verify", txRef))
	flutterReq := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + key},
		},
	}

	result, err := HttpReq(flutterReq)
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

func ServiceTransaction(key, body string) (*ServicedTransaction, error) {
	reqURL, _ := url.Parse(fmt.Sprintf("https://api.flutterwave.com/v3/bills"))
	flutterReq := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + key},
		},
		Body: ioutil.NopCloser(strings.NewReader(body)),
	}

	result, err := HttpReq(flutterReq)
	defer result.Body.Close()
	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		//log.Fatalln(err)
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

func ChargeCard(key, body string) (*FlwVerifiedTransaction, error) {
	reqURL, _ := url.Parse(fmt.Sprintf("https://api.flutterwave.com/v3/tokenized-charges"))
	flutterReq := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + key},
		},
		Body: ioutil.NopCloser(strings.NewReader(body)),
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