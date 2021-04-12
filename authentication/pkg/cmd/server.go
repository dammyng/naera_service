package cmd

import (
	"authentication/internals/config"
	"authentication/pkg"
	"context"
	"log"
	"os"
)

var App pkg.Naera

func RunServers(ctx context.Context) error {

	env := config.NewApConfig()

	err := App.Initialize(env.RedisHost, env.RedisPass, env.AmqpBroker, env.GrpcHost)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	grpcPort := os.Getenv("GRPC_PORT")
	httpPort := os.Getenv("HTTP_PORT")

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
