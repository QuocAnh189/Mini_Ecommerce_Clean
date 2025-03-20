package token

import (
	"strings"
	"time"

	"ecommerce_clean/pkgs/logger"

	"github.com/golang-jwt/jwt"

	"ecommerce_clean/configs"
	"ecommerce_clean/utils"
)

const (
	AccessTokenExpiredTime  = 5 * 60 * 60 // 5 hours
	RefreshTokenExpiredTime = 30 * 24 * 3600
)

type JTWMarker struct {
	AccessTokenType  string
	RefreshTokenType string
}

func NewJTWMarker() (*JTWMarker, error) {
	return &JTWMarker{
		AccessTokenType:  AccessTokenType,
		RefreshTokenType: RefreshTokenType,
	}, nil
}

func (j *JTWMarker) GenerateAccessToken(payload *AuthPayload) string {
	cfg := configs.GetConfig()
	newPayload := NewAuthPayload(payload.ID, payload.Email, payload.Role, time.Minute, AccessTokenType)

	tokenContent := jwt.MapClaims{
		"payload": newPayload,
		"exp":     time.Now().Add(time.Second * AccessTokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate access token: ", err)
		return ""
	}

	return token
}

func (j *JTWMarker) GenerateRefreshToken(payload *AuthPayload) string {
	cfg := configs.GetConfig()
	newPayload := NewAuthPayload(payload.ID, payload.Email, payload.Role, time.Minute, RefreshTokenType)
	tokenContent := jwt.MapClaims{
		"payload": newPayload,
		"exp":     time.Now().Add(time.Second * RefreshTokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate refresh token: ", err)
		return ""
	}

	return token
}

func (j *JTWMarker) ValidateToken(jwtToken string) (*AuthPayload, error) {
	cfg := configs.GetConfig()
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AuthSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	var data *AuthPayload
	utils.MapStruct(&data, tokenData["payload"])

	return data, nil
}
