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

func (n *NaeraBillsRpcServer) UpdateBiller(ctx context.Context, arg *models.UpdateBillerRequest) (*empty.Empty, error) {
	err := n.DB.UpdateABiller(arg.Old, arg.New)
	return &empty.Empty{}, err
}

func (n *NaeraBillsRpcServer) CreateBill(ctx context.Context, arg *models.Bill) (*models.BillCreatedResponse, error) {

	result, err := n.DB.CreateABill(arg)

	if err != nil {
		return nil, err
	}

	return &models.BillCreatedResponse{Id: result}, err
}

func (n *NaeraBillsRpcServer) FindBill(ctx context.Context, arg *models.Bill) (*models.Bill, error) {

	result, err := n.DB.FindABill(arg)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound

	}

	if err != nil {
		return nil, InternalError
	}

	return result, nil
}

func (n *NaeraBillsRpcServer) GetBillerBills(ctx context.Context, arg  *models.GetBillerBillsRequest) (*models.BillsResponse, error) {
	result, err := n.DB.BillerBills(arg.BillerID)
	if err != nil {
		return nil, err
	}
	return &models.BillsResponse{Bills: result}, err
}

func (n *NaeraBillsRpcServer) UpdateBill(ctx context.Context, arg *models.UpdateBillRequest) (*empty.Empty, error) {
	err := n.DB.UpdateABill(arg.Old, arg.New)
	return &empty.Empty{}, err
}

func (n *NaeraBillsRpcServer) CreateBillCategory(ctx context.Context, arg *models.BillCategory) (*models.BillCategoryCreatedResponse, error) {

	result, err := n.DB.CreateABillCategory(arg)

	if err != nil {
		return nil, err
	}

	return &models.BillCategoryCreatedResponse{Id: result}, err
}
func (n *NaeraBillsRpcServer) GetBillCategories(ctx context.Context, arg *empty.Empty) (*models.BillCategoriesResponse, error) {
	result, err := n.DB.BillsCategories()
	if err != nil {
		return nil, err
	}
	return &models.BillCategoriesResponse{Categories: result}, nil
}

func (n *NaeraBillsRpcServer) FindBillCategory(ctx context.Context, arg *models.BillCategory) (*models.BillCategory, error) {

	result, err := n.DB.FindABillCategory(arg)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound

	}

	if err != nil {
		return nil, InternalError
	}

	return result, nil
}

func (n *NaeraBillsRpcServer) UpdateBillCategory(ctx context.Context, arg *models.UpdateBillCategoryRequest) (*empty.Empty, error) {
	err := n.DB.UpdateABillCategory(arg.Old, arg.New)
	return &empty.Empty{}, err
}


func (n *NaeraBillsRpcServer) CreateTransaction(ctx context.Context, arg *models.Transaction) (*models.TransactionCreatedResponse, error) {

	result, err := n.DB.CreateATransaction(arg)

	if err != nil {
		return nil, err
	}

	return &models.TransactionCreatedResponse{Id: result}, err
}


func (n *NaeraBillsRpcServer) BillTransactions(ctx context.Context, arg  *models.GetBillTransactionsRequest) (*models.TransactionsResponse, error) {
	result, err := n.DB.BillTransactions(arg.BillID)
	if err != nil {
		return nil, InternalError
	}
	return &models.TransactionsResponse{Transactions: result}, err
}

func (n *NaeraBillsRpcServer) BillerTransactions(ctx context.Context, arg  *models.GetBillerTransactionsRequest) (*models.TransactionsResponse, error) {
	result, err := n.DB.BillerTransactions(arg.BillerID)
	if err != nil {
		return nil, InternalError
	}
	return &models.TransactionsResponse{Transactions: result}, err
}

func (n *NaeraBillsRpcServer) FindTransaction(ctx context.Context, arg *models.Transaction) (*models.Transaction, error) {

	result, err := n.DB.FindATransaction(arg)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound

	}

	if err != nil {
		return nil, InternalError
	}

	return result, nil
}



func (n *NaeraBillsRpcServer) CreateOrder(ctx context.Context, arg *models.Order) (*models.OrderCreatedResponse, error) {

	result, err := n.DB.CreateAOrder(arg)

	if err != nil {
		return nil, err
	}

	return &models.OrderCreatedResponse{Id: result}, err
}

func (n *NaeraBillsRpcServer) TransactionOrders(ctx context.Context, arg  *models.GetTransactionOrdersRequest) (*models.OrdersResponse, error) {
	result, err := n.DB.TransactionOrders(arg.TransactionID)
	if err != nil {
		return nil, InternalError
	}
	return &models.OrdersResponse{Orders: result}, err
}

func (n *NaeraBillsRpcServer) FindOrder(ctx context.Context, arg *models.Order) (*models.Order, error) {

	result, err := n.DB.FindAOrder(arg)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound

	}

	if err != nil {
		return nil, InternalError
	}

	return result, nil
}



func (n *NaeraBillsRpcServer) CreateCard(ctx context.Context, arg *models.Card) (*models.CardCreatedResponse, error) {

	result, err := n.DB.CreateACard(arg)

	if err != nil {
		return nil, err
	}

	return &models.CardCreatedResponse{Id: result}, err
}

func (n *NaeraBillsRpcServer) GetBillerCards(ctx context.Context, arg  *models.GetBillerCardsRequest) (*models.CardsResponse, error) {
	result, err := n.DB.BillerCards(arg.AddedBy)
	if err != nil {
		return nil, InternalError
	}
	return &models.CardsResponse{Cards: result}, err
}

func (n *NaeraBillsRpcServer) FindCard(ctx context.Context, arg *models.Card) (*models.Card, error) {

	result, err := n.DB.FindACard(arg)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound

	}

	if err != nil {
		return nil, InternalError
	}

	return result, nil
}

func (n *NaeraBillsRpcServer) UpdateCard(ctx context.Context, arg *models.UpdateCardRequest) (*empty.Empty, error) {
	err := n.DB.UpdateACard(arg.Old, arg.New)
	return &empty.Empty{}, err
}


