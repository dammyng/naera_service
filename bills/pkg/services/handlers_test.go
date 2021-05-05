package services_test

import (
	"bills/pkg/services"
	"log"
	"testing"

	"gopkg.in/stretchr/testify.v1/require"
)


func TestVerifiedBill(t *testing.T){
	rs, err := services.FWVerifyBillsHandler("/bill-items/AT099/validate?code=BIL107&customer=09055913141")
	log.Println(rs)
	require.NoError(t, err)
}