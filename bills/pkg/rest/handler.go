package rest

import (
	"bills/models/v1"
	"encoding/json"
	"net/http"
)

type BillHandler struct {
	GrpcPlug models.NaeraBillsServiceClient
}

func NewBillHandler(grpcPlug models.NaeraBillsServiceClient) *BillHandler {
	return &BillHandler{
		GrpcPlug: grpcPlug,
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
