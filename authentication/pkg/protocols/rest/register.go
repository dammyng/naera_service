package rest

import (
	"encoding/json"
	"github.com/twinj/uuid"
	"net/http"
)

func (handler *AuthHandler) AccountRegistration(w http.ResponseWriter, r *http.Request) {
	var reg RegistrationPayload

	err := json.NewDecoder(r.Body).Decode(&reg)
	defer r.Body.Close()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_ = uuid.NewV4().String()

	id, err := handler.DB.CreateUser()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, id)
}
