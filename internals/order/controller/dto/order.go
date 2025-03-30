package dto

import "time"

type Order struct {
	ID         string       `json:"id"`
	Code       string       `json:"code"`
	Lines      []*OrderLine `json:"lines"`
	TotalPrice float64      `json:"total_price"`
	Status     string       `json:"status"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type OrderLine struct {
	Product  Product `json:"product,omitempty"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}

type Product struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
