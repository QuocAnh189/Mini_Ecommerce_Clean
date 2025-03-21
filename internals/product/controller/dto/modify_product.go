package dto

import "mime/multipart"

type CreateProductRequest struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	Price       float64               `form:"price" binding:"gt=0"`
}

type UpdateProductRequest struct {
	Name        string                `form:"name,omitempty"`
	Description string                `form:"description,omitempty"`
	Image       *multipart.FileHeader `form:"image,omitempty"`
	Price       float64               `form:"price,omitempty" binding:"gte=0"`
}
