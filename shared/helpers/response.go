package helpers

import (
	"encoding/json"
	"net/http"
	"github.com/mitchellh/mapstructure"
)

type ResponseData struct {
	Status  string `json:"status" `
	Message string `json:"message" `
	Data    interface{} `json:"data" `
}

func CreateResponse(status, message string, data interface{}) ResponseData {
	var response ResponseData
	response.Status = status
	response.Message = message
	response.Data = data
	return response
}

func (response *ResponseData)DataToStruct(result *interface{})  {
	mapstructure.Decode(response.Data, &result)
}

func RespondWithText(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func RespondWithJSON(w http.ResponseWriter, code int, payload ResponseData) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func SetupCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

	if req.Method == "OPTIONS" {
		(*w).Header().Set("Access-Control-Max-Age", "1728000")

		(*w).Header().Set("Response-Code", "204")
	}

	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}