package pkg

import (
	"bills/internals/db"
	"bills/pkg/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type NaeraBill struct {
	Router *mux.Router
}

func NewNaeraBill() *NaeraBill {
	return &NaeraBill{}
}

func (n *NaeraBill) Initialize() error {

	//
	db := db.NewMockLayer()

	// Router
	router := router.InitServiceRouter(db)
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

func (n *NaeraBill) RunGRPCServer(ctx context.Context, port, dsn string) error {
	return nil
}
