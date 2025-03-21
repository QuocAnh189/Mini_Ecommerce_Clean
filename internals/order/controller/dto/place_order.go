package dto

type PlaceOrderRequest struct {
	UserID string                  `json:"user_id" validate:"required"`
	Lines  []PlaceOrderLineRequest `json:"lines,omitempty" validate:"required,gt=0,lte=5,dive"`
}

type PlaceOrderLineRequest struct {
	ProductID string `json:"product_id,omitempty" validate:"required"`
	Quantity  uint   `json:"quantity,omitempty" validate:"required"`
}
