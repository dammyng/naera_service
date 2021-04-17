package router

import (
	"authentication/pkg/helpers"
	"net/http"
)


func authBearer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.SetupCors(&w, r)
		if r.Method == "OPTIONS"{
			return
		}
		// Do stuff here
		err := helpers.TokenValid(r)
		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden - " + err.Error() , http.StatusForbidden)
		}
	})
}


