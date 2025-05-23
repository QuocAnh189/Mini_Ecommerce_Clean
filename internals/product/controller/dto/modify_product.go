package dto

import "mime/multipart"

type CreateProductRequest struct {
	Name        string                `form:"name" binding:"required" validate:"min=3,max=100"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required" swaggerignore:"true"`
	Price       float64               `form:"price" binding:"gt=0"`
}

type UpdateProductRequest struct {
	ID          string                `form:"id" binding:"required"`
	Name        string                `form:"name,omitempty"`
	Description string                `form:"description,omitempty"`
	Image       *multipart.FileHeader `form:"image,omitempty" swaggerignore:"true"`
	Price       float64               `form:"price,omitempty" binding:"gte=0"`
}
