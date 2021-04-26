package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *AuthHandler) AccountLogin(w http.ResponseWriter, r *http.Request) {
	
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS"{
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	var reg LoginPayload

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

	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusBadRequest, InternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(reg.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Errorf("Invalid username or password").Error())
		return
	}

	
	if u.IsReady == false{
		res := LoginResponse{
			AccessToken:  u.Id,
			RefreshToken: u.Id,
		}
		respondWithJSON(w, http.StatusOK, res)
	return
	}

	ts, err := helpers.CreateToken(u.Id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, InternalServerError)
		return
	}
	saveErr := helpers.CreateAuth(u.Id, ts, handler.RedisService)
	if saveErr != nil {
		respondWithError(w, http.StatusBadRequest, InternalServerError)
	}

	res := LoginResponse{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}
	respondWithJSON(w, http.StatusOK, res)
}
