package cmd

import (
	"bills/config"
	"bills/pkg"
	"context"
	"log"
	"os"
)

var App pkg.NaeraBill

func RunServers(ctx context.Context) error {

	_ = config.NewApConfig()

	err := App.Initialize()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	grpcPort := os.Getenv("GRPC_PORT")
	httpPort := os.Getenv("HTTP_PORT")

	defer cancel()

	
	go func() {
		err = App.RunGRPCServer(ctx, grpcPort, "env.DSN")

		if err != nil {
			log.Panic("GRPC connection failed")
		}
	}()
 
	err = App.RunHTTPServer(ctx, httpPort)
	return err
}
