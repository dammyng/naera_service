package rest

import (
	"authentication/pkg/helpers"
	"encoding/json"
	"net/http"
	"time"

	"github.com/twinj/uuid"
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
		return
	}

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(reg.Email, token, time.Hour)

	//handler.EventEmitter.Emit()

	respondWithJSON(w, http.StatusCreated, id)
}
