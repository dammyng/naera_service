package rest

import (
	"bills/internals/db"
	"encoding/json"
	"net/http"
)

type BillHandler struct {
	DB db.Handler
}

func NewBillHandler(db db.Handler) *BillHandler {
	return &BillHandler{
		DB: db,
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
