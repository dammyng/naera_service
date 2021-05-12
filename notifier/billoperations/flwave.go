package billoperations

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"notifier/models"
	"strings"
	"time"
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


func ServiceTransaction(key, body string) (*models.ServicedTransaction, error) {
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
