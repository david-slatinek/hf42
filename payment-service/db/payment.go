package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"main/model"
	"os"
	"time"
)

type PaymentDB struct {
	database *gorm.DB
}

func NewPaymentDB() (PaymentDB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_HOST")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return PaymentDB{}, err
	}
	return PaymentDB{database: db}, err
}

func (receiver PaymentDB) Close() error {
	db, err := receiver.database.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (receiver PaymentDB) CreateOrder(order model.Order) error {
	payment := model.Payment{
		OrderID:     order.OrderID,
		UserID:      order.CustomerID,
		Amount:      order.TotalPrice,
		PaymentDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	return receiver.database.Create(&payment).Error
}
