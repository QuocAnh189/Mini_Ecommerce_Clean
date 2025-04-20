package token

import (
	"errors"
	"time"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type AuthPayload struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Type      string    `json:"type"`
	Jit       string    `json:"jit"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewAuthPayload(id string, email string, role string, duration time.Duration, type_auth string, jit string) *AuthPayload {
	payload := &AuthPayload{
		ID:        id,
		Email:     email,
		Role:      role,
		Type:      type_auth,
		Jit:       jit,
		ExpiredAt: time.Now().Add(duration),
	}
	return payload
}

func (payload *AuthPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
