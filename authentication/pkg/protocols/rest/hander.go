package rest

import (
	"authentication/internals/db"
	"authentication/myredis"
	"encoding/json"
	"net/http"
	"shared/amqp/sender"
)

type AuthHandler struct {
	DB db.Handler
	RedisService myredis.MyRedis
	EventEmitter sender.EventEmitter
}

func NewAuthHandler(db db.Handler, redis myredis.MyRedis, emitter sender.EventEmitter) *AuthHandler {
	return &AuthHandler{
		DB: db,
		RedisService: redis,
		EventEmitter: emitter,
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