package rest

import (
	"bills/pkg/services"
	"net/http"
)

func (handler *BillHandler) AllDataBundles(w http.ResponseWriter, r *http.Request) {
	res, err := services.GetAllDataBundle()
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}
