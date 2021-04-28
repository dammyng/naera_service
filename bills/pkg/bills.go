package pkg

import (
	"bills/billsgrpc"
	"bills/internals/db"
	"bills/models/migration"
	"bills/models/v1"
	"bills/pkg/router"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type NaeraBill struct {
	Router *mux.Router
}

func NewNaeraBill() *NaeraBill {
	return &NaeraBill{}
}

func (n *NaeraBill) Initialize(grpcHost string) error {

	//GRPC
	grpcClient, err := billsgrpc.NewNaeraRPClient(grpcHost)
	if err != nil {
		return err
	}

	// Router
	router := router.InitServiceRouter(grpcClient)
	n.Router = router
	return nil
}

func (n *NaeraBill) RunHTTPServer(ctx context.Context, port string) error {

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
			log.Panicln(err.Error())
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

func (n *NaeraBill) RunGRPCServer(ctx context.Context, port, dsn string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting GRPC Server on port %v", lis.Addr().String())

	db := db.NewSqlLayer(dsn)
	db.Session.AutoMigrate(migration.Biller{}, migration.Bill{}, migration.BillCategory{})

	grpcServer := grpc.NewServer()
	_naeragrpc := billsgrpc.NewNaeraBillsRpcServer(db)
	models.RegisterNaeraBillingServiceServer(grpcServer, _naeragrpc)
	err = grpcServer.Serve(lis)
	return err}
