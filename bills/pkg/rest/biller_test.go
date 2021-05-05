package rest_test

import (
	"bytes"
	"net/http"
	"os"
	"testing"
	//"github.com/stretchr/testify/require"
)

  func TestUpdateBiller(t *testing.T)  {
	
	var jsonStr = []byte(`{"cart":"[{}]", "card_token":"389" }`)
	req, _ := http.NewRequest("PUT", "/v1/bills/updatebiller", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "bearer " + os.Getenv("test_token"))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	checkResponse(t, http.StatusCreated, response)
	//require.Equal()
}