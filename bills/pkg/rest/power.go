package rest

import (
	"bills/pkg/services"
	"net/http"
)

func (handler *BillHandler) AllPower(w http.ResponseWriter, r *http.Request) {
	res, err := services.GetAllPower()
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}
