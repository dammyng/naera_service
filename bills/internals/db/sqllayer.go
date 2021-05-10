package db

import (
	"bills/models/v1"
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlLayer struct {
	Session *gorm.DB
}

func NewSqlLayer(dsn string) *SqlLayer {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	return &SqlLayer{Session: db}
}

func (sql *SqlLayer) CreateABiller(user *models.Biller) (string, error) {
	err := sql.Session.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.Id, err
}

func (sql *SqlLayer) FindABiller(arg *models.Biller) (*models.Biller, error) {
	session := sql.Session
	var dA models.Biller
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateABiller(old *models.Biller, new *models.Biller) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}

func (sql *SqlLayer) UpdateABillerMap(arg *models.Biller, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}

func (sql *SqlLayer) CreateABill(bill *models.Bill) (string, error) {
	err := sql.Session.Create(&bill).Error
	if err != nil {
		return "", err
	}
	return bill.Id, err
}

func (sql *SqlLayer) BillerBills(arg string) ([]*models.Bill, error) {
	args := models.Bill{Biller: arg}
	session := sql.Session
	 dTs := []*models.Bill{}
	err := session.Where(&args).Find(&dTs).Error
	
	if err != nil {
		return nil, err
	}
	return dTs, nil
}

func (sql *SqlLayer) FindABill(arg *models.Bill) (*models.Bill, error) {
	session := sql.Session
	var dA models.Bill
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateABill(old *models.Bill, new *models.Bill) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}

func (sql *SqlLayer) UpdateABillMap(arg *models.Bill, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}

func (sql *SqlLayer) CreateABillCategory(bCat *models.BillCategory) (string, error) {
	err := sql.Session.Create(&bCat).Error
	if err != nil {
		return "", err
	}
	return bCat.Id, err
}
func (sql *SqlLayer) BillsCategories() ([]*models.BillCategory, error) {
	session := sql.Session
	as := []*models.BillCategory{}
	err := session.Find(&as).Error
	return as, err
}

func (sql *SqlLayer) FindABillCategory(arg *models.BillCategory) (*models.BillCategory, error) {
	session := sql.Session
	var dA models.BillCategory
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateABillCategory(old *models.BillCategory, new *models.BillCategory) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}

func (sql *SqlLayer) UpdateABillCategoryMap(arg *models.BillCategory, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}


func (sql *SqlLayer) CreateATransaction(transaction *models.Transaction) (string, error) {
	err := sql.Session.Create(&transaction).Error
	if err != nil {
		return "", err
	}
	return transaction.Id, err
}

func (sql *SqlLayer) BillerTransactions(arg string) ([]*models.Transaction, error) {
	args := models.Transaction{Biller: arg}
	session := sql.Session
	 dTs := []*models.Transaction{}

	err := session.Where(&args).Find(&dTs).Error
	if err != nil {
		return nil, err
	}
	return dTs, nil
}

func (sql *SqlLayer) BillTransactions(arg string) ([]*models.Transaction, error) {
	args := models.Transaction{Bill: arg}
	session := sql.Session
	 dTs := []*models.Transaction{}

	err := session.Where(&args).Find(&dTs).Error
	if err != nil {
		return nil, err
	}
	return dTs, nil
}

func (sql *SqlLayer) FindATransaction(arg *models.Transaction) (*models.Transaction, error) {
	session := sql.Session
	var dA models.Transaction
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}



func (sql *SqlLayer) CreateAOrder(order *models.Order) (string, error) {
	err := sql.Session.Create(&order).Error
	if err != nil {
		return "", err
	}
	return order.Id, err
}

func (sql *SqlLayer) TransactionOrders(arg string) ([]*models.Order, error) {
	args := models.Order{TransactionId: arg}
	session := sql.Session
	 dTs := []*models.Order{}

	err := session.Where(&args).Find(&dTs).Error
	if err != nil {
		return nil, err
	}
	return dTs, nil
}

func (sql *SqlLayer) FindAOrder(arg *models.Order) (*models.Order, error) {
	session := sql.Session
	var dA models.Order
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}



func (sql *SqlLayer) CreateACard(card *models.Card) (string, error) {
	err := sql.Session.Create(&card).Error
	if err != nil {
		return "", err
	}
	return card.Id, err
}

func (sql *SqlLayer) BillerCards(arg string) ([]*models.Card, error) {
	args := models.Card{AddedBy: arg}
	session := sql.Session
	 dTs := []*models.Card{}

	err := session.Where(&args).Find(&dTs).Error
	if err != nil {
		return nil, err
	}
	return dTs, nil
}

func (sql *SqlLayer) FindACard(arg *models.Card) (*models.Card, error) {
	session := sql.Session
	var dA models.Card
	err := session.Where(arg).First(&dA).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &dA, err
}

func (sql *SqlLayer) UpdateACard(old *models.Card, new *models.Card) error {
	session := sql.Session
	return session.Model(&old).Updates(new).Error
}

func (sql *SqlLayer) UpdateACardMap(arg *models.Card, dict map[string]interface{}) error {
	session := sql.Session
	return session.Model(&arg).Updates(dict).Error
}