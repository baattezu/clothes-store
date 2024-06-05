package cart

type AddToCartDTO struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
