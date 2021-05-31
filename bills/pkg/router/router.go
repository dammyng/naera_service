package router

import (
	"bills/models/v1"
	"bills/pkg/rest"
	"shared/amqp/sender"

	"github.com/gorilla/mux"
)

func InitServiceRouter(grpcPlug models.NaeraBillingServiceClient, emitter sender.EventEmitter) *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewBillHandler(grpcPlug, emitter)

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.Path("/livebills").HandlerFunc(handler.LiveCategories).Methods("GET", "OPTIONS")
	v1.Path("/flbills").HandlerFunc(handler.FLBills).Methods("GET", "OPTIONS")
	v1.Path("/bills/airtime").HandlerFunc(handler.AllAirtimes).Methods("GET", "OPTIONS")
	v1.Path("/bills/cable").HandlerFunc(handler.AllCables).Methods("GET", "OPTIONS")
	v1.Path("/bills/databundle").HandlerFunc(handler.AllDataBundles).Methods("GET", "OPTIONS")
	v1.Path("/bills/internet").HandlerFunc(handler.AllInternet).Methods("GET", "OPTIONS")
	v1.Path("/bills/power").HandlerFunc(handler.AllPower).Methods("GET", "OPTIONS")
	v1.Path("/bills/biller").HandlerFunc(handler.GetBiller).Methods("GET", "OPTIONS")
	v1.Path("/bills/biller/cards").HandlerFunc(handler.BillerCards).Methods("GET", "OPTIONS")
	v1.Path("/bills/updatebiller").HandlerFunc(handler.UpdateBiller).Methods("PUT", "OPTIONS")
	v1.Path("/bills/createbill").HandlerFunc(handler.CreateBill).Methods("POST", "OPTIONS")
	v1.Path("/bills/mybills").HandlerFunc(handler.MyBills).Methods("Get", "OPTIONS")
	v1.Path("/bills/savebill").HandlerFunc(handler.CreateBill).Methods("POST", "OPTIONS")
	v1.Path("/bills/vetnewcart").HandlerFunc(handler.VerifyNewCart).Methods("GET", "OPTIONS")
	v1.Path("/bills/fundWalletfl").HandlerFunc(handler.FundWalletWithFL).Methods("POST", "OPTIONS")

	v1.Path("/bills/{bill_id}").HandlerFunc(handler.GetBill).Methods("GET", "OPTIONS")
	v1.Path("/bill/disable/{bill_id}").HandlerFunc(handler.DisableBill).Methods("PUT", "OPTIONS")
	v1.Path("/bill/delete/{bill_id}").HandlerFunc(handler.DeleteBill).Methods("PUT", "OPTIONS")
	v1.Path("/bills/{bill_id}/transactions").HandlerFunc(handler.BillTransactions).Methods("GET", "OPTIONS")
	v1.Path("/bills/{bill_id}/transaction/{trans_id}").HandlerFunc(handler.BillTransactionOrders).Methods("GET", "OPTIONS")
	//	v1.Path("/bills/paybill/{bill_id}").HandlerFunc(handler.PayForBill).Methods("POST", "OPTIONS")
	v1.Path("/bills/chargewallet").HandlerFunc(handler.ChargeWallet).Methods("POST", "OPTIONS")
	v1.Path("/bills/chargecard").HandlerFunc(handler.ChargeCard).Methods("POST", "OPTIONS")
	v1.Path("/bills/chargeloan").HandlerFunc(handler.ChargeLoan).Methods("POST", "OPTIONS")
	v1.Path("/bills/paywithfl").HandlerFunc(handler.PayWithFL).Methods("POST", "OPTIONS")
	v1.Path("/bills/updatebill/{bill_id}").HandlerFunc(handler.UpdateBill).Methods("PUT", "OPTIONS")
	v1.Path("/biller/transactions").HandlerFunc(handler.BillerTransactions).Methods("GET", "OPTIONS")
	v1.Path("/transaction/{trans_id}").HandlerFunc(handler.BillTransactionOrders).Methods("GET", "OPTIONS")
	v1.Path("/bills/createorder").HandlerFunc(handler.CreateOrder).Methods("POST", "OPTIONS")
	v1.Path("/bills/biller/addcard/{trans_id}").HandlerFunc(handler.AddCard).Methods("POST", "OPTIONS")

	return r
}
