package naeragrpc

import (
	"authentication/internals/db"
	"authentication/models/v1"
	"context"
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

	result, err := n.DB.CreateUser(arg)

	if err != nil {
		return nil, InternalError
	}

	return &models.UserCreatedResponse{Id: result}, nil
}
