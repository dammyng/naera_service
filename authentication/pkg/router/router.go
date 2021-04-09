package router

import (
	"authentication/internals/db"
	"authentication/myredis"
	"authentication/pkg/protocols/rest"

	"github.com/gorilla/mux"
	"shared/amqp/sender"
)

func InitServiceRouter(db db.Handler, redis myredis.MyRedis, emitter sender.EventEmitter) *mux.Router {
	var r = mux.NewRouter()
	handler := rest.NewAuthHandler(db, redis, emitter)

	r.Methods("GET", "POST").Path("/").HandlerFunc(handler.LiveCheck)

	//v1
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/login", handler.AccountLogin).Methods("POST")
	v1.HandleFunc("/register", handler.AccountRegistration).Methods("POST")
	v1.HandleFunc("/verify/{email}/{token}", handler.VerifyEmail).Methods("GET")
	v1.Use(authBearer)
	v1.HandleFunc("/sendverification", handler.SendVerification).Methods("POST")
	v1.HandleFunc("/newpassword", handler.NewPassword).Methods("POST")
	v1.HandleFunc("/ressetpassword", handler.ResetPasssword).Methods("POST")

	return r
}
