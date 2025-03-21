package http

import (
	"ecommerce_clean/internals/order/controller/dto"
	"ecommerce_clean/internals/order/usecase"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/response"
	"ecommerce_clean/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	usecase usecase.IOrderUseCase
}

func NewOrderHandler(usecase usecase.IOrderUseCase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

// PlaceOrder godoc
//
//	@Summary	place order
//	@Tags		orders
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		_	body		dto.PlaceOrderReq	true	"Body"
//	@Success	200	{object}	dto.Order
//	@Router		/api/v1/orders [post]
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

// GetOrders godoc
//
//	@Summary	get my orders
//	@Tags		orders
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		_	query		dto.ListOrderReq	true	"Query"
//	@Success	200	{object}	dto.ListOrderRes
//	@Router		/api/v1/orders [get]
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

// GetOrderByID godoc
//
//	@Summary	get order details
//	@Tags		orders
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id	path		string	true	"Order ID"
//	@Success	200	{object}	dto.Order
//	@Router		/api/v1/orders/{id} [get]
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

// UpdateOrder godoc
//
//	@Summary	cancel order
//	@Tags		orders
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id	path	string	true	"Order ID"
//	@Router		/api/v1/orders/{id}/{status} [put]
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
