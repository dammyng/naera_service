package router

import (
	"bills/models/v1"
	"bills/pkg/rest"

	"github.com/gorilla/mux"
)




func InitServiceRouter(grpcPlug models.NaeraBillsServiceClient) *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewBillHandler(grpcPlug)

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.Methods("GET").Path("/livebills").HandlerFunc(handler.LiveCategories)
	v1.Methods("GET").Path("/bills/airtime").HandlerFunc(handler.AllAirtimes)
	v1.Methods("GET").Path("/bills/cable").HandlerFunc(handler.AllCables)
	v1.Methods("GET").Path("/bills/databundle").HandlerFunc(handler.AllDataBundles)
	v1.Methods("GET").Path("/bills/internet").HandlerFunc(handler.AllInternet)
	v1.Methods("GET").Path("/bills/power").HandlerFunc(handler.AllPower)
	v1.Methods("PUT").Path("/bills/updatebiller").HandlerFunc(handler.UpdateBiller)

	return r	
}