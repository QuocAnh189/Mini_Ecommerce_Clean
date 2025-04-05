package http

import (
	"ecommerce_clean/db"
	"ecommerce_clean/internals/cart/usecase"
	"ecommerce_clean/pkgs/middlewares"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
	"ecommerce_clean/pkgs/validation"

	"github.com/gin-gonic/gin"

	cartRepo "ecommerce_clean/internals/cart/repository"
	productRepo "ecommerce_clean/internals/product/repository"
)

func Routes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	cache redis.IRedis,
	token token.IMarker,
) {

	cartRepository := cartRepo.NewCartRepository(sqlDB)
	productRepository := productRepo.NewProductRepository(sqlDB)
	cartUseCase := usecase.NewCartUseCase(validator, cartRepository, productRepository)
	cartHandler := NewCartHandler(cartUseCase)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	cartRoute := r.Group("/carts", authMiddleware)
	{
		cartRoute.GET("/:userID", cartHandler.GetCart)
		cartRoute.POST("/:userID", cartHandler.AddProductToCart)
		cartRoute.PUT("/cart-line/:userID", cartHandler.UpdateCartLine)
		cartRoute.DELETE("/:userID", cartHandler.RemoveProductToCart)
	}
}
