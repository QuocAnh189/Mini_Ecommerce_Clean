package dto

type Product struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
