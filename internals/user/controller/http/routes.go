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
	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signin", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "SignIn Route"})
		})
		authRouter.POST("/signup", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "SignUp Route"})
		})
		authRouter.POST("/signout", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "SignIn Route"})
		})
	}

	userRouter := r.Group("/users")
	{
		userRouter.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Get Users"})
		})
		userRouter.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Get User"})
		})
		userRouter.DELETE("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"test": "Delete User"})
		})
	}
}
