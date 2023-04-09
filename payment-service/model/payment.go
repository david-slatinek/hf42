package model

type Payment struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	OrderID     string  `json:"orderID"`
	UserID      string  `json:"userID"`
	Amount      float32 `json:"amount"`
	PaymentDate string  `json:"paymentDate"`
}
