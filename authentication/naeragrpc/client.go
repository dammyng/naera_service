package naeragrpc

import (
	"authentication/models/v1"
	"google.golang.org/grpc"
)

type NaeraRPClient struct {
	Conn *grpc.ClientConn
}

func NewNaeraRPClient(addr string) (models.NaeraServiceClient, error) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return models.NewNaeraServiceClient(conn), nil

}
