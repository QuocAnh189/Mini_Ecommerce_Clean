package dto

import (
	"ecommerce_clean/internals/user/entity"
	"mime/multipart"
)

type SignUpRequest struct {
	Email    string                `form:"email" binding:"required,email"`
	Name     string                `form:"name" binding:"required"`
	Avatar   *multipart.FileHeader `form:"avatar"`
	Password string                `form:"password" binding:"required"`
}

type SignUpResponse struct {
	AccessToken  string       `json:"accessToken" validate:"required"`
	RefreshToken string       `json:"refreshToken" validate:"required"`
	User         *entity.User `json:"user" validate:"required"`
}
