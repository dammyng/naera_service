package router

import (
	"bills/models/v1"
	"bills/pkg/rest"

	"github.com/gorilla/mux"
)




func InitServiceRouter(grpcPlug models.NaeraBillingServiceClient) *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewBillHandler(grpcPlug)

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.Path("/livebills").HandlerFunc(handler.LiveCategories).Methods("GET", "OPTIONS")
	v1.Path("/bills/airtime").HandlerFunc(handler.AllAirtimes).Methods("GET", "OPTIONS")
	v1.Path("/bills/cable").HandlerFunc(handler.AllCables).Methods("GET", "OPTIONS")
	v1.Path("/bills/databundle").HandlerFunc(handler.AllDataBundles).Methods("GET", "OPTIONS")
	v1.Path("/bills/internet").HandlerFunc(handler.AllInternet).Methods("GET", "OPTIONS")
	v1.Path("/bills/power").HandlerFunc(handler.AllPower).Methods("GET", "OPTIONS")
	v1.Path("/bills/updatebiller").HandlerFunc(handler.UpdateBiller).Methods("PUT", "OPTIONS")
	v1.Path("/bills/createbill").HandlerFunc(handler.CreateBill).Methods("POST", "OPTIONS")
	v1.Path("/bills/updatebill/{bill_id}").HandlerFunc(handler.UpdateBill).Methods("PUT", "OPTIONS")
	return r	
}