package naeragrpc

import (
	"authentication/models/v1"
	"context"

	"google.golang.org/grpc"
)

type naeraServiceClient struct {
	Conn *grpc.ClientConn
}

func NewNaeraRPClient(addr string) (*naeraServiceClient, error) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	//opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	return &naeraServiceClient{
		Conn: conn,
	}, nil
}

func (c *naeraServiceClient) RegisterAccount(ctx context.Context, in *models.Account, opts ...grpc.CallOption) (*models.UserCreatedResponse, error) {
	ss := models.NewNaeraServiceClient(c.Conn)
	result, err := ss.RegisterAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}
