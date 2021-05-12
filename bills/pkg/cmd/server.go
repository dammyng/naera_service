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

	env := config.NewApConfig()

	err := App.Initialize(env.GrpcHost, env.AmqpBroker)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	grpcPort := os.Getenv("BILL_GRPC_PORT")
	httpPort := os.Getenv("BILL_HTTP_PORT")

	defer cancel()

	
	go func() {
		err = App.RunGRPCServer(ctx, grpcPort, env.DSN)

		if err != nil {
			log.Panic("GRPC connection failed")
		}
	}()
 
	err = App.RunHTTPServer(ctx, httpPort)
	return err
}
