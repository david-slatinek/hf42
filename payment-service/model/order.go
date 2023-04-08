package model

import "time"

type Order struct {
	CustomerID string    `json:"customer_id"`
	OrderDate  time.Time `json:"order_date"`
	Books      []Book    `json:"books"`
	TotalPrice float32   `json:"total_price"`
	Status     string    `json:"status"`
}
