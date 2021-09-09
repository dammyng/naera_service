package router

import (
	"authentication/pkg/helpers"
	"net/http"
	"strings"
)

func authBearer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.SetupCors(&w, r)
		if r.Method == "OPTIONS" {
			return
		}

		bearToken := r.Header.Get("Authorization")
		//normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Do stuff here
		err := helpers.TokenValid(r)
		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	})
}
