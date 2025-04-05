package http

import (
	"ecommerce_clean/internals/order/controller/dto"
	"ecommerce_clean/internals/order/usecase"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/response"
	"ecommerce_clean/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase usecase.IOrderUseCase
}

func NewOrderHandler(usecase usecase.IOrderUseCase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

// @Summary			Place a new order
// @Description		Creates a new order for the authenticated user.
// @Tags			Orders
// @Produce			json
// @Security		ApiKeyAuth
// @Param			request	body	dto.PlaceOrderRequest	true	"Order details"
// @Success			200	{object}	dto.Order	"Order placed successfully"
// @Failure			400	{object}	response.Response	"Bad Request - Invalid parameters"
// @Failure			401	{object}	response.Response	"Unauthorized - User not authenticated"
// @Failure			403	{object}	response.Response	"Forbidden - User does not have the required permissions"
// @Failure			500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
// @Router			/orders [post]
// @Security		ApiKeyAuth
func (a *OrderHandler) PlaceOrder(c *gin.Context) {
	var req dto.PlaceOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	req.UserID = c.GetString("userId")
	if req.UserID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	order, err := a.usecase.PlaceOrder(c, &req)
	if err != nil {
		logger.Error("Failed to create OrderHandler: ", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.Order
	utils.MapStruct(&res, &order)
	response.JSON(c, http.StatusOK, res)
}

// @Summary			Get my orders
// @Description		Retrieve a list of orders for the authenticated user with optional filters.
// @Tags			Orders
// @Produce			json
// @Security		ApiKeyAuth
// @Param			code		query	string	false	"Filter by order code"
// @Param			status		query	string	false	"Filter by order status"
// @Param			page		query	int		false	"Page number for pagination (default: 1)"
// @Param			limit		query	int		false	"Number of records per page (default: 10)"
// @Param			order_by	query	string	false	"Field to order by (e.g., created_at)"
// @Param			order_desc	query	bool	false	"Sort order: true for descending, false for ascending"
// @Success			200	{object}	dto.ListOrdersResponse	"Orders retrieved successfully"
// @Failure			400	{object}	response.Response		"Bad Request - Invalid parameters"
// @Failure			401	{object}	response.Response		"Unauthorized - User not authenticated"
// @Failure			500	{object}	response.Response		"Internal Server Error - An error occurred while processing the request"
// @Router			/orders [get]
// @Security		ApiKeyAuth
func (a *OrderHandler) GetOrders(c *gin.Context) {
	var req dto.ListOrdersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Failed to parse request req: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	req.UserID = c.GetString("userId")
	if req.UserID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	orders, pagination, err := a.usecase.ListMyOrders(c, &req)
	if err != nil {
		logger.Error("Failed to get orders: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.ListOrdersResponse
	res.Pagination = pagination
	utils.MapStruct(&res.Orders, &orders)
	response.JSON(c, http.StatusOK, res)
}

// @Summary			Get order details
// @Description		Retrieve details of a specific order by its ID.
// @Tags			Orders
// @Produce			json
// @Security		ApiKeyAuth
// @Param			id	path		string	true	"Order ID"
// @Success			200	{object}	dto.Order		"Order retrieved successfully"
// @Failure			400	{object}	response.Response	"Bad Request - Missing or invalid Order ID"
// @Failure			401	{object}	response.Response	"Unauthorized - User not authenticated"
// @Failure			404	{object}	response.Response	"Not Found - Order does not exist"
// @Failure			500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
// @Router			/orders/{id} [get]
// @Security		ApiKeyAuth
func (a *OrderHandler) GetOrderByID(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	orderId := c.Param("id")
	if orderId == "" {
		response.Error(c, http.StatusBadRequest, errors.New("bad request"), "Miss Order ID")
		return
	}

	order, err := a.usecase.GetOrderByID(c, orderId)
	if err != nil {
		logger.Errorf("Failed to get order, id: %s, error: %s ", orderId, err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.Order
	utils.MapStruct(&res, &order)
	response.JSON(c, http.StatusOK, res)
}

// @Summary			Update order status
// @Description		Update the status of an existing order.
// @Tags			Orders
// @Produce			json
// @Security		ApiKeyAuth
// @Param			id		path	string	true	"Order ID"
// @Param			status	path	string	true	"New order status"
// @Success			200	{object}	dto.Order		"Order updated successfully"
// @Failure			400	{object}	response.Response	"Bad Request - Missing or invalid Order ID"
// @Failure			401	{object}	response.Response	"Unauthorized - User not authenticated"
// @Failure			404	{object}	response.Response	"Not Found - Order does not exist"
// @Failure			500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
// @Router			/orders/{id}/{status} [put]
// @Security		ApiKeyAuth
func (a *OrderHandler) UpdateOrder(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	orderID := c.Param("id")
	if orderID == "" {
		response.Error(c, http.StatusBadRequest, errors.New("bad request"), "Miss Order ID")
		return
	}

	status := c.Param("status")
	order, err := a.usecase.UpdateOrder(c, orderID, userID, status)
	if err != nil {
		logger.Errorf("Failed to cancel order, id: %s, error: %s", orderID, err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.Order
	utils.MapStruct(&res, &order)
	response.JSON(c, http.StatusOK, res)
}
