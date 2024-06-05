package order

type CreateOrderDTO struct {
	UserID string      `json:"user_id"`
	Items  []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type Order struct {
	OrderID string      `json:"order_id"`
	UserID  string      `json:"user_id"`
	Items   []OrderItem `json:"items"`
}
