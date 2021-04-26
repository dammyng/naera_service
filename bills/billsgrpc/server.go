package billsgrpc

import (
	"bills/internals/db"
	"bills/models/v1"
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"gorm.io/gorm"
)


type NaeraBillsRpcServer struct {
	DB db.Handler
}


func NewNaeraBillsRpcServer(db db.Handler) *NaeraBillsRpcServer {
	
	return &NaeraBillsRpcServer{
		DB: db,
	}
}

func (n *NaeraBillsRpcServer) CreateBiller(ctx context.Context, arg *models.Biller) (*models.BillerCreatedResponse, error) {

	result, err := n.DB.CreateABiller(arg)

	if err != nil {
		return nil, err
	}

	return &models.BillerCreatedResponse{Id: result}, err
}

func (n *NaeraBillsRpcServer) FindBiller(ctx context.Context, arg *models.Biller) (*models.Biller, error) {

	result, err := n.DB.FindABiller(arg)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound

	}
	
	if err != nil {
		return nil, InternalError
	}

	return result, nil
}

func (n *NaeraBillsRpcServer) UpdateBiller(ctx context.Context, arg *models.UpdateBillerRequest) (*empty.Empty , error) {
	err := n.DB.UpdateABiller(arg.Old, arg.New)
	return &empty.Empty{},err
}