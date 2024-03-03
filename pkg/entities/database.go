package entities

type Order struct {
	ID        string `json:"id"`
	UserID    string `json:"userID"`
	OrderDate string `json:"orderDate"`
	Status    string `json:"status"`
}

type OrderItem struct {
	ID           string
	OrderID      string
	ItemID       string
	Quantity     int
	PricePerUnit float64
}
