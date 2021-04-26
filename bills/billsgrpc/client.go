package billsgrpc

import (
	"bills/models/v1"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type naeraBillsServiceClient struct {
	Conn *grpc.ClientConn
}


func NewNaeraRPClient(addr string) (*naeraBillsServiceClient, error) {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	//opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	return &naeraBillsServiceClient{
		Conn: conn,
	}, nil
}

func (c *naeraBillsServiceClient) CreateBiller(ctx context.Context, in *models.Biller, opts ...grpc.CallOption) (*models.BillerCreatedResponse, error) {
	ss := models.NewNaeraBillsServiceClient(c.Conn)
	result, err := ss.CreateBiller(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) FindBiller(ctx context.Context, in *models.Biller, opts ...grpc.CallOption) (*models.Biller, error) {
	ss := models.NewNaeraBillsServiceClient(c.Conn)
	result, err := ss.FindBiller(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) UpdateBiller(ctx context.Context, in *models.UpdateBillerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.Conn.Invoke(ctx, "/models.NaeraService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}