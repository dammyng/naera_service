package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"net/http"

	"google.golang.org/grpc"
)

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
