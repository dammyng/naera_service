package rest

import (
	"net"
	"net/http"
	"time"
)


var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}

var NetClient = &http.Client{
	Timeout:   time.Second * 10,
	Transport: netTransport,
}
