package middlewares

import (
	"encoding/json"
	"fmt"
	"ecommerce_clean/pkgs/logger"
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
)

type AuthMiddleware struct {
	token token.IMarker
}

func NewAuthMiddleware(token token.IMarker) *AuthMiddleware {
	return &AuthMiddleware{
		token: token,
	}
}

func (a *AuthMiddleware) TokenAuth(cache redis.IRedis) gin.HandlerFunc {
	return a.Token(token.AccessTokenType, cache)
}

func (a *AuthMiddleware) TokenRefresh(cache redis.IRedis) gin.HandlerFunc {
	return a.Token(token.RefreshTokenType, cache)
}

func (a *AuthMiddleware) Token(tokenType string, cache redis.IRedis) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		payload, err := a.token.ValidateToken(token)
		if err != nil || payload == nil || payload.Type != tokenType {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		// Lấy dữ liệu từ Redis
		var rawValue string
		if err := cache.Get(fmt.Sprintf("blacklist:%s_%s", payload.ID, payload.Jit), &rawValue); err != nil {
			logger.Error("Failed to get value from Redis:", err)
		}

		var value map[string]string
		err = json.Unmarshal([]byte(rawValue), &value)
		if err != nil {
			logger.Error("Failed to unmarshal JSON:", err)
		}

		if value["status"] == "blacklisted" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			c.Abort()
			return
		}

		c.Set("userId", payload.ID)
		c.Set("role", payload.Role)
		c.Set("jit", payload.Jit)
		c.Set("token", token)
		c.Next()
	}
}
