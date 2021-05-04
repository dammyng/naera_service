package db

import "bills/models/v1"


type Handler interface {
	CreateABiller(*models.Biller) (string, error)
	FindABiller(*models.Biller) (*models.Biller, error)
	UpdateABiller(*models.Biller, *models.Biller) error
	UpdateABillerMap(*models.Biller, map[string]interface{}) error

	BillerBills(string) ([]*models.Bill, error)
	CreateABill(*models.Bill) (string, error)
	FindABill(*models.Bill) (*models.Bill, error)
	UpdateABill(*models.Bill, *models.Bill) error


	BillsCategories() ([]*models.BillCategory, error)
	CreateABillCategory(*models.BillCategory) (string, error)
	FindABillCategory(*models.BillCategory) (*models.BillCategory, error)
	UpdateABillCategory(*models.BillCategory, *models.BillCategory) error

	CreateATransaction(*models.Transaction) (string, error)
	BillerTransactions(string) ([]*models.Transaction, error)
	FindATransaction(*models.Transaction) (*models.Transaction, error)



	CreateAOrder(*models.Order) (string, error)
	TransactionOrders(string) ([]*models.Order, error)
	FindAOrder(*models.Order) (*models.Order, error)


}
