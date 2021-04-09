package rest

import (
	"authentication/internals/db"
	"authentication/myredis"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	DB db.Handler
	RedisService myredis.MyRedis
}

func NewAuthHandler(db db.Handler,  redis myredis.MyRedis) *AuthHandler {
	return &AuthHandler{
		DB: db,
		RedisService: redis,
	}
}

func (handler *AuthHandler) LiveCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (handler *AuthHandler) LiveUpdate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}