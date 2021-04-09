package rest

import (
	"authentication/pkg/helpers"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"shared/amqp/events"
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

	id := uuid.NewV4()

	_, err = handler.DB.CreateUser()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, id)

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(reg.Email, token, time.Hour)

	msg := events.UserCreatedEvent{
		ID:    hex.EncodeToString(id.Bytes()),
		Email: reg.Email,
		Token: token,
	}

	handler.EventEmitter.Emit(&msg, "NaeraAuth")

}
