package rest

import (
	"encoding/json"
	"net/http")



type BillHandler struct {

}

func NewBillHandler() *BillHandler {
	return &BillHandler{
	
	}
}

func (handler *BillHandler) LiveCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}