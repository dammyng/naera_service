package rest

import (
	"bills/pkg/services"
	"net/http"
)

func (handler *BillHandler) AllInternet(w http.ResponseWriter, r *http.Request) {
	res, err := services.GetAllAirtime()
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}
