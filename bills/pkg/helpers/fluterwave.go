package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

