package helpers

import (
	"authentication/models/v1"
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

func AccountReady(a *models.Account)  bool {
	/**if a.EmailVerifiedAt < 1000  || a.PhoneVerifiedAt < 1000 ||  a.BvnVerifiedAt < 1000 ||  a.NubanVerifiedAt < 1000 {
		return false
	}
	if a.IdCard == "" {
		return false
	}
	**/
	if a.EmailVerifiedAt < 1000 {
		return false
	}
	if a.IdCard == "" {
		return false
	}
	return true
}