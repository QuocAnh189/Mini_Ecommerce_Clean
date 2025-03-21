package dto

import (
	"ecommerce_clean/internals/user/entity"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	AccessToken  string       `json:"accessToken" validate:"required"`
	RefreshToken string       `json:"refreshToken" validate:"required"`
	User         *entity.User `json:"user" validate:"required"`
}
