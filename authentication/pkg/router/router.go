package router

import (
	"authentication/models/v1"
	"authentication/myredis"
	"authentication/pkg/protocols/rest"

	"shared/amqp/sender"

	"github.com/gorilla/mux"
)

func InitServiceRouter(redis myredis.MyRedis, emitter sender.EventEmitter, grpcPlug models.NaeraServiceClient) *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewAuthHandler(redis, emitter, grpcPlug)

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)

	//v1
	v1 := r.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/v1/login", handler.AccountLogin).Methods("POST", "OPTIONS")
	r.HandleFunc("/v1/register", handler.AccountRegistration).Methods("POST", "OPTIONS")
	r.HandleFunc("/v1/verify/{email}/{token}", handler.VerifyEmail).Methods("GET", "OPTIONS")
	r.HandleFunc("/v1/newpassword/{email}", handler.NewPassword).Methods("POST", "OPTIONS")
	r.HandleFunc("/v1/resetpassword", handler.ResetPasssword).Methods("POST", "OPTIONS")
	v1.Use(authBearer)
	v1.HandleFunc("/sendverification/{email}", handler.SendVerification).Methods("POST", "OPTIONS")

	return r
}
