package router

import (
	"authentication/pkg/helpers"
	"net/http")


func authBearer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		err := helpers.TokenValid(r)
		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden - " + err.Error() , http.StatusForbidden)
		}
	})
}
