package http

import (
	"ecommerce_clean/db"
	"ecommerce_clean/pkgs/minio"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
	"ecommerce_clean/pkgs/validation"
	"github.com/gin-gonic/gin"
)

func Routes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	minioClient *minio.MinioClient,
	cache redis.IRedis,
	token token.IMarker,
) {
	orderRoute := r.Group("/orders")
	{
		orderRoute.POST("", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Place Order"})
		})
		orderRoute.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Get Orders"})
		})
		orderRoute.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Get Order"})
		})
		orderRoute.PUT("/:id/cancel", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Cancel Order"})
		})
	}
}
