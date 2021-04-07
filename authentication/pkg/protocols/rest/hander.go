package rest

import (
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (handler *AuthHandler) LiveCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}
