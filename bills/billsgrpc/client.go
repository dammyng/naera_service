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
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	out, err := ss.GetBillerBills(ctx, in, opts...)
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


func (c *naeraBillsServiceClient) CreateTransaction(ctx context.Context, in *models.Transaction, opts ...grpc.CallOption) (*models.TransactionCreatedResponse, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.CreateTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) BillerTransactions(ctx context.Context, in *models.GetBillerTransactionsRequest, opts ...grpc.CallOption) (*models.TransactionsResponse, error) {
	var out *models.TransactionsResponse
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	out, err := ss.BillerTransactions(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naeraBillsServiceClient) BillTransactions(ctx context.Context, in *models.GetBillTransactionsRequest, opts ...grpc.CallOption) (*models.TransactionsResponse, error) {
	
	var out *models.TransactionsResponse
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	out, err := ss.BillTransactions(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


func (c *naeraBillsServiceClient) FindTransaction(ctx context.Context, in *models.Transaction, opts ...grpc.CallOption) (*models.Transaction, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.FindTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}


func (c *naeraBillsServiceClient) CreateOrder(ctx context.Context, in *models.Order, opts ...grpc.CallOption) (*models.OrderCreatedResponse, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.CreateOrder(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) TransactionOrders(ctx context.Context, in *models.GetTransactionOrdersRequest, opts ...grpc.CallOption) (*models.OrdersResponse, error) {
	var out *models.OrdersResponse
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	out, err := ss.TransactionOrders(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


func (c *naeraBillsServiceClient) FindOrder(ctx context.Context, in *models.Order, opts ...grpc.CallOption) (*models.Order, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.FindOrder(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}



func (c *naeraBillsServiceClient) CreateCard(ctx context.Context, in *models.Card, opts ...grpc.CallOption) (*models.CardCreatedResponse, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.CreateCard(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *naeraBillsServiceClient) GetBillerCards(ctx context.Context, in *models.GetBillerCardsRequest, opts ...grpc.CallOption) (*models.CardsResponse, error) {
	var out *models.CardsResponse
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	out, err := ss.GetBillerCards(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


func (c *naeraBillsServiceClient) FindCard(ctx context.Context, in *models.Card, opts ...grpc.CallOption) (*models.Card, error) {
	ss := models.NewNaeraBillingServiceClient(c.Conn)
	result, err := ss.FindCard(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return result, err
}


func (c *naeraBillsServiceClient) UpdateCard(ctx context.Context, in *models.UpdateCardRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.Conn.Invoke(ctx, "/models.NaeraBillingService/UpdateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}