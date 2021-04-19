package main

import (
	"bills/pkg/cmd"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	err := cmd.RunServers(ctx)
	if err != nil {
		log.Fatalf("Error Starting servers %v", err.Error())
	}
}
