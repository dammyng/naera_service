package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"
	"log"

	"encoding/hex"
	"encoding/json"
	"net/http"

	"shared/amqp/events"
	"time"

	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func (handler *AuthHandler) AccountRegistration(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	var reg RegistrationPayload

	err := json.NewDecoder(r.Body).Decode(&reg)
	defer r.Body.Close()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	var opts []grpc.CallOption

	u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Email: reg.Email}, opts...)

	if u != nil {
		respondWithError(w, http.StatusBadRequest, DuplicateUserAccount)
		return
	}

	id := uuid.NewV4()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	
	res, err := handler.GrpcPlug.RegisterAccount(r.Context(), &models.Account{Id: id.String(), Email: reg.Email, FirstName: reg.FirstName, Surname: reg.LastName, PhoneNumber: reg.Phone, Password: hashedPass, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}, opts...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, res.Id)

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(reg.Email, token, time.Hour)
	log.Println(token)
	msg := events.UserCreatedEvent{
		ID:    hex.EncodeToString(id.Bytes()),
		Email: reg.Email,
		Token: token,
	}

	handler.EventEmitter.Emit(&msg, "NaeraExchange")
	
}
