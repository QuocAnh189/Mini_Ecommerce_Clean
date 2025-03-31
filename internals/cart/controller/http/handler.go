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

//	@Summary		Retrieve the cart of a user
//	@Description	Fetches the shopping cart details of the authenticated user based on the provided user ID.
//	@Tags			Carts
//	@Produce		json
//	@Param			userID	path	string	true	"User ID"
//	@Success		200	{object}	dto.Cart	"Successfully retrieved the user's cart"
//	@Failure		400	{object}	response.Response	"Bad Request - Invalid request parameters"
//	@Failure		401	{object}	response.Response	"Unauthorized - User ID mismatch or authentication failed"
//	@Failure		403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure		404	{object}	response.Response	"Not Found - Cart not found for the given user ID"
//	@Failure		500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router			/carts/{userID} [get]
//	@Security		ApiKeyAuth
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

//	@Summary		Add a product to the user's cart
//	@Description	Adds a specified product to the authenticated user's shopping cart.
//	@Tags			Carts
//	@Accept			json
//	@Produce		json
//	@Param			userID		path	string					true	"User ID"
//	@Param			body		body	dto.AddProductRequest	true	"Product details to add to cart"
//	@Success		201			{string}	string				"Add product to cart successfully"
//	@Failure		400			{object}	response.Response	"Bad Request - Invalid request parameters"
//	@Failure		401			{object}	response.Response	"Unauthorized - User ID mismatch or authentication failed"
//	@Failure		403			{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure		500			{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router			/carts/{userID} [post]
//	@Security		ApiKeyAuth
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

//	@Summary		Update a cart line item
//	@Description	Updates the quantity or details of a specific product in the authenticated user's shopping cart.
//	@Tags			Carts
//	@Accept			json
//	@Produce		json
//	@Param			userID		path	string						true	"User ID"
//	@Param			body		body	dto.UpdateCartLineRequest	true	"Updated cart line details"
//	@Success		200			{string}	string					"Update cart successfully"
//	@Failure		400			{object}	response.Response		"Bad Request - Invalid request parameters"
//	@Failure		401			{object}	response.Response		"Unauthorized - User ID mismatch or authentication failed"
//	@Failure		403			{object}	response.Response		"Forbidden - User does not have the required permissions"
//	@Failure		500			{object}	response.Response		"Internal Server Error - An error occurred while processing the request"
//	@Router			/carts/cart-line/{userID} [put]
//	@Security		ApiKeyAuth
func (h *CartHandler) UpdateCartLine(c *gin.Context) {
	userID := c.GetString("userId")
	userIDParam := c.Param("userID")

	if userID == "" || userIDParam == "" || userID != userIDParam {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	var req dto.UpdateCartLineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	if err := h.usecase.UpdateCartLine(c, &req); err != nil {
		logger.Error("Failed to update cart", err)
		return
	}

	response.JSON(c, http.StatusCreated, "Update cart successfully")
}

//	@Summary		Remove a product from the user's cart
//	@Description	Removes a specified product from the authenticated user's shopping cart.
//	@Tags			Carts
//	@Accept			json
//	@Produce		json
//	@Param			userID		path	string						true	"User ID"
//	@Param			body		body	dto.RemoveProductRequest	true	"Product details to remove from cart"
//	@Success		200			{string}	string					"Remove product from cart successfully"
//	@Failure		400			{object}	response.Response		"Bad Request - Invalid request parameters"
//	@Failure		401			{object}	response.Response		"Unauthorized - User ID mismatch or authentication failed"
//	@Failure		403			{object}	response.Response		"Forbidden - User does not have the required permissions"
//	@Failure		500			{object}	response.Response		"Internal Server Error - An error occurred while processing the request"
//	@Router			/carts/{userID} [delete]
//	@Security		ApiKeyAuth
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
