package dto

type Order struct {
	ID         string       `json:"id"`
	Code       string       `json:"code"`
	Lines      []*OrderLine `json:"lines"`
	TotalPrice float64      `json:"total_price"`
	Status     string       `json:"status"`
}

type OrderLine struct {
	Product  Product `json:"product,omitempty"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}

type Product struct {
	ID    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
