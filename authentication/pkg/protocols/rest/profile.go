package rest

import (
	"authentication/models/migration"
	"authentication/models/v1"
	"authentication/pkg/helpers"
	"encoding/json"
	"log"
	"time"

	//"encoding/hex"

	"net/http"

	//"shared/amqp/events"

	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	tokenAuth, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	userId, err := helpers.FetchAuth(tokenAuth, handler.RedisService)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	var opts []grpc.CallOption

	user, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Id: userId}, opts...)
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (handler *AuthHandler) FindProfiles(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	_, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	params := mux.Vars(r)
	query := params["query"]
	var opts []grpc.CallOption

	users, err := handler.GrpcPlug.FindAccounts(r.Context(), &models.FindAccountsRequest{Query: query}, opts...)
	if err != nil {

		respondWithError(w, http.StatusInternalServerError, grpc.ErrorDesc(err))
		return
	}

	if len(users.Accounts) == 0 {
		respondWithJSON(w, http.StatusOK, make([]string, 0))
		return
	}
	respondWithJSON(w, http.StatusOK, users.Accounts)
}

func (handler *AuthHandler) GetSetUpProfile(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)

	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}
	key := helpers.ExtractToken(r)

	var opts []grpc.CallOption

	user, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Id: key}, opts...)
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
	var cleanUser migration.CleanAccount
	copier.Copy(&cleanUser, &user)
	respondWithJSON(w, http.StatusOK, cleanUser)
}

func (handler *AuthHandler) UpdateSetUpProfile(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}
	key,err := helpers.ExtractTokenMetadata(r)

	var opts []grpc.CallOption

	var u models.Account
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if len(u.Password) > 0 {
		respondWithError(w, http.StatusOK, "You cannot reset your password here!")
	}
	user, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Id: key.UserId}, opts...)
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
	if len(u.Pin) > 1 {
		hashedPin, err := bcrypt.GenerateFromPassword([]byte(u.Pin), bcrypt.DefaultCost)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
		u.PinUpdatedAt = time.Now().Unix()
		u.Pin = hashedPin
	}
	_, err = handler.GrpcPlug.UpdateAccount(r.Context(), &models.UpdateAccountRequest{Old: user, New: &u}, opts...)
	if err != nil {
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}

	user, err = handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Id: key.UserId}, opts...)
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}

	profileIsReady := helpers.AccountReady(user)
	log.Println(profileIsReady)

	_, err = handler.GrpcPlug.UpdateAccount(r.Context(), &models.UpdateAccountRequest{Old: user, New: &models.Account{IsReady: profileIsReady, UpdatedAt: time.Now().Unix()}}, opts...)
	if err != nil {
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
	var cleanUser migration.CleanAccount
	copier.Copy(&cleanUser, &user)
	respondWithJSON(w, http.StatusOK, cleanUser)
}

func (handler *AuthHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	tokenAuth, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	userId, err := helpers.FetchAuth(tokenAuth, handler.RedisService)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	var opts []grpc.CallOption

	var u models.Account
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	user, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Id: userId}, opts...)
	isReady := user.IsReady
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
	_, err = handler.GrpcPlug.UpdateAccount(r.Context(), &models.UpdateAccountRequest{Old: user, New: &u}, opts...)
	if err != nil {
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}
	profileIsReady := helpers.AccountReady(user)
	if profileIsReady != isReady {
		_, err = handler.GrpcPlug.UpdateAccount(r.Context(), &models.UpdateAccountRequest{Old: user, New: &models.Account{IsReady: profileIsReady}}, opts...)
		if err != nil {
			respondWithError(w, http.StatusOK, InternalServerError)
			return
		}
	}

	var cleanUser migration.CleanAccount
	copier.Copy(&cleanUser, &user)
	respondWithJSON(w, http.StatusOK, cleanUser)
}
