package pkg

import (
	"authentication/internals/db"
	"authentication/myredis"
	"authentication/pkg/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Naera struct {
	Router *mux.Router
}

func NewNaera() *Naera {
	return &Naera{}
}

// Initialize service components
// 1. Database Handler
// 2. Reddis client
// 3. Logger
// 4. GRPC Client
// 5. RabbitMQ Emitter
// 6. Router
func (n *Naera) Initialize(dsn, redisHost, redisPass string,) error {

	//DB
	db := db.NewSqlLayer(dsn)

	//Redis
	redis := myredis.NewMyRedis(redisHost, redisPass)

	// Router
	router := router.InitServiceRouter(db, redis)
	n.Router = router
	return nil
}

// RunHTTPServer starts a http server for the application
func (n *Naera) RunHTTPServer(ctx context.Context, port string) error {

	server := &http.Server{
		Addr:           port,
		Handler:        n.Router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Starting HTTP Server on port %v", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Panicln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// until the timeout deadline.
	server.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
	return nil
}

// RunGRPCServer starts a GRPC server for the application
func (n *Naera) RunGRPCServer(ctx context.Context, port string) error {
	log.Println("Starting GRPC Server")

	return nil
}
