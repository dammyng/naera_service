package pkg

import (
	"authentication/internals/db"
	models "authentication/models/v1"
	"authentication/myredis"
	"authentication/naeragrpc"
	"authentication/pkg/router"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"shared/amqp/sender"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
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
func (n *Naera) Initialize(redisHost, redisPass, amqpHost, grpcHost string) error {

	//GRPC
	grpcClient, err := naeragrpc.NewNaeraRPClient(grpcHost)
	if err != nil {
		return  err
	}
	//Redis
	redis := myredis.NewMyRedis(redisHost, redisPass)

	//AMQP
	conn, err := amqp.Dial(amqpHost)
	if err != nil {
		return err
	}
	eventEmitter, err := sender.NewAmqpEventEmitter(conn, "NaeraAuth")
	if err != nil {
		return err
	}

	// Router
	router := router.InitServiceRouter(redis, eventEmitter, grpcClient)
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
func (n *Naera) RunGRPCServer(ctx context.Context, port , dsn string) error {
	log.Println("Starting GRPC Server")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := db.NewSqlLayer(dsn)

	grpcServer := grpc.NewServer()
	_naeragrpc := naeragrpc.NewNaeraRpcServer(db)
	models.RegisterNaeraServiceServer(grpcServer, _naeragrpc)
	err = grpcServer.Serve(lis)
	return err
}
