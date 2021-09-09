package billoperations

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"notifier/models"
	"strings"
	"time"
	"os"
	"io"
)


func HttpReq(req *http.Request) (*http.Response, error) {

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 15 * time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 15,
		Transport: netTransport,
	}
	res, err := netClient.Do(req)

	// check for response error
	if err != nil {
		log.Fatal("Error:", err)
		return nil, err
	}
	return res, err
}


func ServiceTransaction(body string) (*models.ServicedTransaction, error) {	
	
	flutterReq, err := BuildFlutterWaveRequest("POST", "/bills", ioutil.NopCloser(strings.NewReader(body)))
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
		//log.Fatalln(err)
		return nil, err
	}
	var response models.ServicedTransaction
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status == "error" {
		return nil, errors.New(response.Message)
	}
	return &response, err
}


func BuildFlutterWaveRequest(method, path string, body io.Reader) (*http.Request, error) {
	BASE_URL := "https://api.flutterwave.com/v3"

	req, err := http.NewRequest(method, BASE_URL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if os.Getenv("Environment") == "production" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("FL_SECRETKEY_LIVE")))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("FL_SECRETKEY_TEST")))
	}
	return req, nil
}