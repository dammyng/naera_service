package pkg

import (
	"context"
	"log"
	"naerarauth/internals/persistence"
	"naerarauth/memstore"
	"naerarauth/pkg/router/v1"
	"naerarshared/interfaces"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type NaerarAuthHandler struct {
	Router *mux.Router
	Db     persistence.NaerarAuthDBHandler
	Memstore interfaces.MemStorage
}

func (e *NaerarAuthHandler) InitDb(db persistence.NaerarAuthDBHandler) {
	e.Db = db
}

func (e *NaerarAuthHandler) InitMemStore(cache memstore.Redis) {
	e.Memstore = cache
}

func (e *NaerarAuthHandler) InitRouter() {
	rr := router.InitRoutes(e.Db, e.Memstore)
	e.Router = rr
}

func (e *NaerarAuthHandler) StartHttp(port string) {
	server := &http.Server{
		Addr:           port,
		Handler:        e.Router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Run our server in a goroutine so that it doesn't block.
	log.Printf("Saterting server on port: %v", port)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
