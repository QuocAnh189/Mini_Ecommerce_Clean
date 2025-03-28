package http

import (
	"ecommerce_clean/internals/cart/controller/dto"
	"ecommerce_clean/internals/cart/usecase"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/response"
	"ecommerce_clean/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartHandler struct {
	usecase usecase.ICartUseCase
}

func NewCartHandler(usecase usecase.ICartUseCase) *CartHandler {
	return &CartHandler{
		usecase: usecase,
	}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.GetString("userId")
	userIDParam := c.Param("userID")

	if userID == "" || userIDParam == "" || userID != userIDParam {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	cart, err := h.usecase.GetCartByUserID(c, userID)
	if err != nil {
		logger.Errorf("Failed to get cart by user, id: %s, error: %s ", userID, err)
		response.Error(c, http.StatusNotFound, err, "Not found")
	}

	var res *dto.Cart
	utils.MapStruct(&res, cart)
	response.JSON(c, http.StatusOK, res)
}

func (h *CartHandler) AddProductToCart(c *gin.Context) {
	userID := c.GetString("userId")
	userIDParam := c.Param("userID")

	if userID == "" || userIDParam == "" || userID != userIDParam {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	var req dto.AddProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	if err := h.usecase.AddProduct(c, &req); err != nil {
		logger.Error("Failed to add product to cart", err)
		return
	}

	response.JSON(c, http.StatusCreated, "Add product to cart successfully")
}

func (h *CartHandler) RemoveProductToCart(c *gin.Context) {
	userID := c.GetString("userId")
	userIDParam := c.Param("userID")

	if userID == "" || userIDParam == "" || userID != userIDParam {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	var req dto.RemoveProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	if err := h.usecase.RemoveProduct(c, &req); err != nil {
		logger.Error("Failed to add product", err)
		return
	}

	response.JSON(c, http.StatusOK, "Remove product from cart successfully")
}
