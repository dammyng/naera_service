package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"bills/pkg/restclient"
	"encoding/json"
	"net/http"
	"time"

	"github.com/twinj/uuid"
)

func (handler *BillHandler) FundWalletWithFL(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var fundPayload restclient.FundWalletPayload
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&fundPayload)
	if err != nil {
		http.Error(w,  err.Error(), 400)
		return
	}

	verified, err := restclient.VerifyFwTransaction(fundPayload.TransactionID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	newTransaction := models.Transaction{
		Id:            uuid.NewV4().String(),
		Class:         "credit",
		Biller:        access.UserId,
		Title:         "Wallet Fund",
		WalletID:      fundPayload.WalletID,
		BillingMethod: "",
		Amount:        float32(verified.Data.Amount),
		TransRef:      fundPayload.TransactionID,
		Bill:          fundPayload.WalletID,
		FlRef:         verified.Data.FewRef,
		CreatedAt:     time.Now().Unix(),
	}

	tran_id, err := handler.GrpcPlug.CreateTransaction(r.Context(), &newTransaction)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, tran_id)
}
