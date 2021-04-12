package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"shared/amqp/events"
	"time"

	"github.com/twinj/uuid"
	"google.golang.org/grpc"
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

	var opts []grpc.CallOption
	_, err = handler.GrpcPlug.RegisterAccount(r.Context(), &models.Account{Id: id.String(), Email: reg.Email, FirstName: reg.FirstName, Surname: reg.LastName, PhoneNumber: reg.Phone, }, opts...)
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
