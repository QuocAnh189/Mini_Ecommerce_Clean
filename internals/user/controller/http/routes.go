package http

import (
	"ecommerce_clean/db"
	"ecommerce_clean/internals/user/repository"
	"ecommerce_clean/internals/user/usecase"
	"ecommerce_clean/pkgs/middlewares"
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
	userRepository := repository.NewUserRepository(sqlDB)
	userUseCase := usecase.NewUserUseCase(validator, userRepository, minioClient, cache, token)
	userHandler := NewAuthHandler(userUseCase)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signup", userHandler.SignUp)
		authRouter.POST("/signin", userHandler.SignIn)
		authRouter.POST("/signout", authMiddleware, userHandler.SignOut)
	}

	userRouter := r.Group("/users").Use(authMiddleware)
	{
		userRouter.GET("", userHandler.GetUsers)
		userRouter.GET("/:id", userHandler.GetUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
	}
}
