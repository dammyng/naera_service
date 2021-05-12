package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"google.golang.org/grpc"
)

func (handler *BillHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	var u models.Order
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if strArr[1] != "CreWwdvOO3pclP3ZFqZUbsDZYL0HyWoU"{
		respondWithError(w, http.StatusUnauthorized, errors.New("INVALID API KEY").Error())
		return
	}

	var opts []grpc.CallOption

	order := &models.Order{
		Id: u.Id,
		Title: u.Title,
		TransactionId: u.TransactionId,
		Amount: u.Amount,
		Charged: u.Charged,
		Fulfilled: u.Charged,
		CreatedAt: u.CreatedAt,
	}

	tRes, err := handler.GrpcPlug.CreateOrder(r.Context(), order, opts...)
	if err != nil {
		err = errors.New("Error creating the bill record")
		respondWithError(w, http.StatusBadRequest, err.Error())

	}
	respondWithJSON(w, http.StatusCreated, tRes.Id)
		
		
}
