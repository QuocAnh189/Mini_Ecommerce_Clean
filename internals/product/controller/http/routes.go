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
	productRoute := r.Group("/products")
	{
		productRoute.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Get Products"})
		})
		productRoute.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Get Product"})
		})
		productRoute.POST("", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Create Product"})
		})
		productRoute.PUT("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Update Product"})
		})
		productRoute.DELETE("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Delete Product"})
		})
	}
}
