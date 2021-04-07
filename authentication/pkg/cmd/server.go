package cmd

import (
	"authentication/pkg"
	"context"
	"os"
)

var App pkg.Naera

func RunServers(ctx context.Context) error {
	err := App.Initialize()

	ctx, cancel := context.WithCancel(ctx)
	grpcPort := os.Getenv("GRPC_PORT")
	httpPort := os.Getenv("HTTP_PORT")

	defer cancel()

	go func() {
		err = App.RunGRPCServer(ctx, grpcPort)

	}()

	err = App.RunHTTPServer(ctx, httpPort)
	return err
}
