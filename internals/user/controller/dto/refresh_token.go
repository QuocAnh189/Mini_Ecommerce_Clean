package dto

type RefreshTokenReq struct {
	RefreshToken string `json:"refreshToken" binding:"required" validate:"required"`
}

type RefreshTokenRes struct {
	AccessToken string `json:"accessToken"`
}
