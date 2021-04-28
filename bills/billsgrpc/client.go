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
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.CreateBiller(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) FindBiller(ctx context.Context, in *models.Biller, opts ...grpc.CallOption) (*models.Biller, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.FindBiller(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) UpdateBiller(ctx context.Context, in *models.UpdateBillerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.Conn.Invoke(ctx, "/models.NaeraBillingService/UpdateBiller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naeraBillsServiceClient) CreateBill(ctx context.Context, in *models.Bill, opts ...grpc.CallOption) (*models.BillCreatedResponse, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.CreateBill(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) GetBillerBills(ctx context.Context, in *models.GetBillerBillsRequest, opts ...grpc.CallOption) (*models.BillsResponse, error) {
	var out *models.BillsResponse
	err := c.Conn.Invoke(ctx, "/models.NaeraBillingService/GetBillerBills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naeraBillsServiceClient) FindBill(ctx context.Context, in *models.Bill, opts ...grpc.CallOption) (*models.Bill, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.FindBill(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) UpdateBill(ctx context.Context, in *models.UpdateBillRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.Conn.Invoke(ctx, "/models.NaeraBillingService/UpdateBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naeraBillsServiceClient) CreateBillCategory(ctx context.Context, in *models.BillCategory, opts ...grpc.CallOption) (*models.BillCategoryCreatedResponse, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.CreateBillCategory(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) GetBillCategories(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*models.BillCategoriesResponse, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	res, err := ss.GetBillCategories(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *naeraBillsServiceClient) FindBillCategory(ctx context.Context, in *models.BillCategory, opts ...grpc.CallOption) (*models.BillCategory, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.FindBillCategory(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) UpdateBillCategory(ctx context.Context, in *models.UpdateBillCategoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.Conn.Invoke(ctx, "/models.NaeraBillingService/UpdateBillCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
