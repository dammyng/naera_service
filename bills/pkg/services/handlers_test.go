package services_test

import (
	"bills/pkg/services"
	"log"
	"testing"

	"gopkg.in/stretchr/testify.v1/require"
	m "bills/models"

)


func TestVerifiedBill(t *testing.T){
	_order := m.OrderRequest{
		ItemCode: "AT099",
		Customer: "08069475323",
		BillerCode: "BIL099",
	}
	rs, err := services.ValidateOrderItem(_order)
	log.Println(rs.Data.Name)
	require.NoError(t, err)
}