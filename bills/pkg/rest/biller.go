package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *BillHandler) CreateBiller(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	var u models.Biller
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	u.Id = access.UserId

	var opts []grpc.CallOption

	biller := &models.Biller{Id: access.UserId,
		CardToken: u.CardToken,
		Cart: u.Cart,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	res, err := handler.GrpcPlug.CreateBiller(r.Context(), biller, opts...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, grpc.ErrorDesc(err))
		return
	}
	respondWithJSON(w, http.StatusCreated, res.Id)

}

func (handler *BillHandler) UpdateBiller(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	var u models.Biller
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var opts []grpc.CallOption
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	u.Id = access.UserId
	biller, err := handler.GrpcPlug.FindBiller(r.Context(), &models.Biller{Id: access.UserId}, opts...)

	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			_, err := handler.GrpcPlug.CreateBiller(r.Context(), &u, opts...)
			if err != nil {
				respondWithError(w, http.StatusNotFound, BillerNotFound)
			}
			respondWithJSON(w, http.StatusCreated, "")
			return
		}
		respondWithError(w, http.StatusBadRequest, err.Error() + "---"+ InternalServerError)
		return
	}

	_, err = handler.GrpcPlug.UpdateBiller(r.Context(), &models.UpdateBillerRequest{Old: biller, New: &u}, opts...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	respondWithJSON(w, http.StatusCreated, "")

}
