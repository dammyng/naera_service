package rest_test

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestBillerCards(t *testing.T) {

	req, _ := http.NewRequest("GET", "/v1/bills/biller/cards", nil)
	req.Header.Set("Authorization", "bearer "+os.Getenv("test_token"))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	log.Println(response.Body)
	checkResponse(t, http.StatusOK, response)
}
