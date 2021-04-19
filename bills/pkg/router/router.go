package router

import (
	"bills/pkg/rest"
	"github.com/gorilla/mux"
)




func InitServiceRouter() *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewBillHandler()

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)

	return r	
}