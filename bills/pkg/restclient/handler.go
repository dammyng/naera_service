package restclient

import (
	"log"
	"net"
	"net/http"
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