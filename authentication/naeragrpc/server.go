package naeragrpc

import (
	"authentication/internals/db"
	"authentication/models/v1"
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"gorm.io/gorm"
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
		return nil, err
	}

	return &models.UserCreatedResponse{Id: result}, err
}



func (n *NaeraRpcServer) FindAccount(ctx context.Context, arg *models.Account) (*models.Account, error) {

	result, err := n.DB.FindUser(arg)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound

	}
	
	if err != nil {
		return nil, InternalError
	}

	return result, nil
}

func (n *NaeraRpcServer) UpdateAccount(ctx context.Context, arg *models.UpdateAccountRequest) (*empty.Empty , error) {
	err := n.DB.UpdateUser(arg.Old, arg.New)
	return &empty.Empty{},err
}