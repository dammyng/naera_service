package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *BillHandler) UpdateBiller(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS"{
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

	biller, err := handler.GrpcPlug.FindBiller()

	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusBadRequest, InternalServerError)
		return
	}

	_, err = handler.GrpcPlug.UpdateBiller(r.Context(), &models.UpdateBillerRequest{Old: biller, New: &u}, opts...)
	if err != nil {
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
}
