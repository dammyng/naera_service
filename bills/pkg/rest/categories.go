package rest

import (
	"net/http"
)

func (handler *BillHandler) LiveCategories(w http.ResponseWriter, r *http.Request) {
	//res, err := handler.DB.GetLiveCategories()
	//if err != nil {
	//	respondWithError(w, http.StatusBadRequest, err.Error())
	//}
	respondWithJSON(w, http.StatusCreated, nil)
}
