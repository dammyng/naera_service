package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"bills/pkg/restclient"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (handler *BillHandler) AddCard(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}
	var opts []grpc.CallOption
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	params := mux.Vars(r)
	trans_id := params["trans_id"]


	verified, err := restclient.VerifyFwTransaction(string(trans_id))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	newCard := models.Card{
		Id: uuid.NewV4().String(),
		Token: verified.Data.Card.Token,
		Email: verified.Data.Customer.Email,
		Status: "active",
		LastDigits: verified.Data.Card.Last4Digits,
		FirstDigits: verified.Data.Card.First6Digits,
		Provider: "FL",
		Expires: verified.Data.Card.Expiry,
		AddedBy: access.UserId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	_, err = handler.GrpcPlug.FindCard(r.Context(), &models.Card{Provider: newCard.Provider, FirstDigits: newCard.FirstDigits, LastDigits: newCard.LastDigits, AddedBy: newCard.AddedBy})
	if status.Convert(err).Message() == gorm.ErrRecordNotFound.Error(){
		res, err := handler.GrpcPlug.CreateCard(r.Context(), &newCard, opts...)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, res.Id)
	}else{
		respondWithJSON(w, http.StatusOK, "")
	}
}

func (handler *BillHandler) BillerCards(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}
	var opts []grpc.CallOption
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	res, err := handler.GrpcPlug.GetBillerCards(r.Context(), &models.GetBillerCardsRequest{AddedBy: access.UserId}, opts...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(res.Cards) ==0 {
		respondWithJSON(w, http.StatusOK, make([]string, 0))
		return
	}
	respondWithJSON(w, http.StatusOK, res.Cards)

}
