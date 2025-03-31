package dto

type Cart struct {
	ID    string      `json:"id"`
	User  *User       `json:"user"`
	Lines []*CartLine `json:"lines"`
}

type CartLine struct {
	ID       string   `json:"id"`
	Product  *Product `json:"product"`
	Quantity int64    `json:"quantity"`
	Price    float64  `json:"price"`
}

type AddProductRequest struct {
	CartID    string `json:"cart_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type UpdateCartLineRequest struct {
	ID        string `json:"id" validate:"required"`
	CartID    string `json:"cart_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type RemoveProductRequest struct {
	CartID    string `json:"cart_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
}
