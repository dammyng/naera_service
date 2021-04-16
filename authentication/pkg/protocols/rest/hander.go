package rest

import (
	"authentication/models/v1"
	"authentication/myredis"
	"encoding/json"
	"net/http"
	"shared/amqp/sender"
)

type AuthHandler struct {
	RedisService myredis.MyRedis
	EventEmitter sender.EventEmitter
	GrpcPlug     models.NaeraServiceClient
}

func NewAuthHandler(redis myredis.MyRedis, emitter sender.EventEmitter, grpcPlug models.NaeraServiceClient) *AuthHandler {
	return &AuthHandler{
		RedisService: redis,
		EventEmitter: emitter,
		GrpcPlug:     grpcPlug,
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

func setupCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

	if req.Method == "OPTIONS" {
		(*w).Header().Set("Access-Control-Max-Age", "1728000")
		(*w).Header().Set("Response-Code", "204")
	}

	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
