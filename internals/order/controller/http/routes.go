package http

import (
	"ecommerce_clean/db"
	"ecommerce_clean/internals/order/repository"
	"ecommerce_clean/internals/order/usecase"
	productRepo "ecommerce_clean/internals/product/repository"
	"ecommerce_clean/pkgs/middlewares"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
	"ecommerce_clean/pkgs/validation"

	"github.com/gin-gonic/gin"
)

func Routes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	cache redis.IRedis,
	token token.IMarker,
) {
	productRepository := productRepo.NewProductRepository(sqlDB)
	orderRepository := repository.NewOrderRepository(sqlDB)
	orderUsecase := usecase.NewOrderUseCase(validator, orderRepository, productRepository)
	orderHandler := NewOrderHandler(orderUsecase)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	orderRoute := r.Group("/orders", authMiddleware)
	{
		orderRoute.POST("", orderHandler.PlaceOrder)
		orderRoute.GET("", orderHandler.GetOrders)
		orderRoute.GET("/:id", orderHandler.GetOrderByID)
		orderRoute.PUT("/:id/:status", orderHandler.UpdateOrder)
	}
}
