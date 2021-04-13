package rest

import (
	"net/http"
	"github.com/gorilla/mux"
)


func (handler *AuthHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	
	//Get route parameters
	params := mux.Vars(r)
	email := params["email"]
	reqToken := params["token"]
	err := handler.RedisService.Client.Get("email")

}


func (handler *AuthHandler) SendVerification(w http.ResponseWriter, r *http.Request) {

}
