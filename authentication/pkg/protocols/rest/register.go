package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"
	"log"
	"os"

	"encoding/hex"
	"encoding/json"
	"net/http"

	"shared/amqp/events"
	"time"

	valid "github.com/asaskevich/govalidator"

	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func (handler *AuthHandler) AccountRegistration(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	var reg RegistrationPayload

	err := json.NewDecoder(r.Body).Decode(&reg)
	defer r.Body.Close()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = valid.ValidateStruct(reg)
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

	res, err := handler.GrpcPlug.RegisterAccount(r.Context(), &models.Account{Id: id.String(), Email: reg.Email, FirstName: reg.FirstName, Surname: reg.LastName, PhoneNumber: reg.Phone, Password: hashedPass, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix(), IsReady: false, WalletID: helpers.RandInt(10)}, opts...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, grpc.ErrorDesc(err))
		return
	}

	_res := LoginResponse{
		AccessToken:  res.Id,
		RefreshToken: res.Id,
	}
	respondWithJSON(w, http.StatusCreated, _res)

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(reg.Email, token, time.Hour)
	log.Println(token)

	if os.Getenv("Environment") == "production" {
		msg := events.UserCreatedEvent{
			ID:    hex.EncodeToString(id.Bytes()),
			Email: reg.Email,
			Token: token,
		}

		handler.EventEmitter.Emit(&msg, "NaeraExchange")
	}

}
