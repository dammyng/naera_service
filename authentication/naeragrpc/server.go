package naeragrpc

import (
	"authentication/internals/db"
	"authentication/models/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NaeraRpcServer struct {
	DB db.Handler
}

func NewNaeraRpcServer(db db.Handler) *NaeraRpcServer {
	return &NaeraRpcServer{
		DB: db,
	}
}


func (n *NaeraRpcServer) RegisterAccount(ctx context.Context, arg *models.Account) (*models.UserCreatedResponse, error) {
    return nil, status.Error(codes.Unimplemented, "Execute() not implemented yet")
}