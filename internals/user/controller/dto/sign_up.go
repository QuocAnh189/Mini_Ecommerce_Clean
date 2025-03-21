package dto

import (
	"ecommerce_clean/internals/user/entity"
	"mime/multipart"
)

type SignUpRequest struct {
	Email    string                `form:"email" validate:"required,email"`
	Name     string                `form:"name" validate:"required"`
	Avatar   *multipart.FileHeader `form:"avatar"`
	Password string                `form:"password" validate:"required"`
}

type SignUpResponse struct {
	AccessToken  string       `json:"accessToken" validate:"required"`
	RefreshToken string       `json:"refreshToken" validate:"required"`
	User         *entity.User `json:"user" validate:"required"`
}
