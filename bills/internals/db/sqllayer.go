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
	args := models.Bill{Id: arg}
	session := sql.Session
	 dTs := []*models.Bill{}

	err := session.Where(&args).Find(dTs).Error
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
