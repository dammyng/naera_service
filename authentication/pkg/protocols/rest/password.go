package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"
	"log"

	//"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	//"shared/amqp/events"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *AuthHandler) NewPassword(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)
	params := mux.Vars(r)
	email := params["email"]
	var opts []grpc.CallOption

	u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Email: email}, opts...)
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, fmt.Errorf("No user was found with the email address:  %v", email).Error())
			return
		}
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("You would receive a mail soon at %v", u.Email)})

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(fmt.Sprintf("%s_password_reset", u.Email), token, time.Hour)
	log.Println(token)
	/*msg := events.PasswordResetRequest{
		ID:    hex.EncodeToString([]byte(u.Id)),
		Email: email,
		Token: token,
	}

	handler.EventEmitter.Emit(&msg, "NaeraExchange")*/

}

func (handler *AuthHandler) ResetPasssword(w http.ResponseWriter, r *http.Request) {

	var reg ResetPasswordPayload

	err := json.NewDecoder(r.Body).Decode(&reg)
	defer r.Body.Close()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	var opts []grpc.CallOption

	u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Email: reg.Email}, opts...)
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, fmt.Errorf("No user was found with the email address:  %v", u.Email).Error())
			return
		}
		respondWithError(w, http.StatusBadRequest, InternalServerError)
		return
	}

	storedToken, err := handler.RedisService.Client.Get(fmt.Sprintf("%s_password_reset", reg.Email)).Result()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("Invalid or incorrect token").Error())
		return
	}
	if storedToken == "" {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("Invalid or incorrect token").Error())
		return
	}
	if match := reg.Token == storedToken; match {

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		_, err = handler.GrpcPlug.UpdateAccount(r.Context(), &models.UpdateAccountRequest{Old: u, New: &models.Account{Password: hashedPass}}, opts...)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, InternalServerError)
			return
		}
		handler.RedisService.Client.Del(fmt.Sprintf("%s_password_reset", reg.Email))
		respondWithJSON(w, http.StatusOK, map[string]interface{}{"message": PassswordResetSuccessful})
		return
	} else {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("Invalid or incorrect token").Error())
		return
	}
}
