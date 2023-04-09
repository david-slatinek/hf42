package model

type Order struct {
	OrderID    string  `json:"orderID"`
	CustomerID string  `json:"customerID"`
	OrderDate  string  `json:"orderDate"`
	Books      []Book  `json:"books"`
	TotalPrice float32 `json:"totalPrice"`
	Status     string  `json:"status"`
}
