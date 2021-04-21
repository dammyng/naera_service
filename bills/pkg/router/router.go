package router

import (
	"bills/internals/db"
	"bills/pkg/rest"

	"github.com/gorilla/mux"
)




func InitServiceRouter(db db.Handler) *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewBillHandler(db)

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)
	r.Methods("GET").Path("/livebills").HandlerFunc(handler.LiveCategories)
	r.Methods("GET").Path("/bills/airtime").HandlerFunc(handler.AllAirtimes)
	r.Methods("GET").Path("/bills/cable").HandlerFunc(handler.AllCables)
	r.Methods("GET").Path("/bills/databundle").HandlerFunc(handler.AllDataBundles)
	r.Methods("GET").Path("/bills/internet").HandlerFunc(handler.AllInternet)
	r.Methods("GET").Path("/bills/power").HandlerFunc(handler.AllPower)

	return r	
}