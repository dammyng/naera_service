package helpers

import (
	"net/http"
)

func SetupCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method == "OPTIONS" {
		(*w).Header().Set("Access-Control-Max-Age", "1728000")
		(*w).Header().Set("Response-Code", "204")
	}

	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
