package dto

import "mime/multipart"

type CreateProductRequest struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required" swaggerignore:"true"`
	Price       float64               `form:"price" binding:"gt=0"`
}

type UpdateProductRequest struct {
	Name        string                `form:"name,omitempty"`
	Description string                `form:"description,omitempty"`
	Image       *multipart.FileHeader `form:"image,omitempty" swaggerignore:"true"`
	Price       float64               `form:"price,omitempty" binding:"gte=0"`
}
