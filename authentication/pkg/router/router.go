package router

import (
	"authentication/pkg/protocols/rest"

	"github.com/gorilla/mux"
)

func InitServiceRouter() *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewAuthHandler()

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)

	return r
}
