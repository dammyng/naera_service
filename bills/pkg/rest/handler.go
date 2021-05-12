package rest

import (
	"bills/models/v1"
	"encoding/json"
	"net/http"
	"shared/amqp/sender"

)

type BillHandler struct {
	GrpcPlug models.NaeraBillingServiceClient
	EventEmitter sender.EventEmitter

}

func NewBillHandler(grpcPlug models.NaeraBillingServiceClient, emitter sender.EventEmitter) *BillHandler {
	return &BillHandler{
		GrpcPlug: grpcPlug,
		EventEmitter: emitter,
	}
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

func (handler *BillHandler) LiveCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
