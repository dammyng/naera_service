package router

import (
	"naerarauth/internals/persistence"
	handlers "naerarauth/pkg/handlers/v1"
	"naerarshared/interfaces"

	"github.com/gorilla/mux"
)

func InitRoutes(db persistence.NaerarAuthDBHandler, redis interfaces.MemStorage) *mux.Router {
	var r = mux.NewRouter()

	v1Handlers := handlers.NewNaerarAuthRouteHandler(db, redis)
	r.HandleFunc("/v1/login", v1Handlers.LoginAccount).Methods("POST", "OPTIONS")
	r.HandleFunc("/v1/register", v1Handlers.CreateAccount).Methods("POST", "OPTIONS")

	v1routes := r.PathPrefix("/v1").Subrouter()
	v1routes.HandleFunc("/me", v1Handlers.MeProfile).Methods("GET", "OPTIONS")

	return r
}
