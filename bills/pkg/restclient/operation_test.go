package restclient_test

import (
	"bills/pkg/restclient"
	"fmt"
	"log"
	"testing"

	"gopkg.in/stretchr/testify.v1/require"
)


func TestFlwTransacVerify(t *testing.T)  {
	r, err := restclient.VerifyFwTransaction("FLWSECK_TEST-be6475503d295c1be0b10ee8e971671f-X", "2062317")
	log.Println(r)
	require.NoError(t, err)
}


func TestFlwServiceTransaction(t *testing.T)  {
	payload := `{
		"country": "NG",
		"customer": "+23408069475323",
		"amount": 100,
		"recurrence": "ONCE",
		"type": "AIRTIME",
		"reference": "9300049404445"
	 }`
	 _, err := restclient.ServiceTransaction("FLWSECK_TEST-be6475503d295c1be0b10ee8e971671f-X", payload)
	require.NoError(t, err)
}

func TestFlwChargingCard(t *testing.T)  {
	payload := `{
		"token": "flw-t1nf-bb9d62ecec403546b7ffe85b58de2ffe-m03k",
		"currency": "NGN",
		"country": "NG",
		"amount": 200,
		"email": "dammy@gmail.com",
		"first_name": "Dami",
		"last_name": "Kassim",
		"narration": "Sample tokenized charge",
		"tx_ref": "tokenid-c-001"
	}`
	_res, result := restclient.ChargeCard("FLWSECK_TEST-be6475503d295c1be0b10ee8e971671f-X", payload)
	log.Println(_res)
	fmt.Println(_res)
	require.NoError(t, result)
}