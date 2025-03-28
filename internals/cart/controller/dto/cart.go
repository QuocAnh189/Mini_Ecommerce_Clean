package dto

type Cart struct {
	ID    string             `json:"id"`
	User  *User              `json:"user"`
	Lines []*CartLineRequest `json:"lines"`
}

type CartLine struct {
	Product  *Product `json:"product"`
	Quantity uint     `json:"quantity" validate:"required"`
}

type CartLineRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  uint   `json:"quantity" validate:"required"`
}

type AddProductRequest struct {
	CartID    string `json:"cart_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type RemoveProductRequest struct {
	CartID    string `json:"cart_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
}
