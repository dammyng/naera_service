package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *AuthHandler) AccountLogin(w http.ResponseWriter, r *http.Request) {
	var reg LoginPayload

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
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusBadRequest, InternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(reg.Password))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("Invalid username or password").Error())
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
