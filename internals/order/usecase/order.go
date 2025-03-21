package usecase

import (
	"context"
	"ecommerce_clean/internals/order/controller/dto"
	"ecommerce_clean/internals/order/entity"
	"ecommerce_clean/internals/order/repository"
	productEntity "ecommerce_clean/internals/product/entity"
	productRepo "ecommerce_clean/internals/product/repository"
	"ecommerce_clean/pkgs/paging"
	"ecommerce_clean/pkgs/validation"
	"ecommerce_clean/utils"
	"errors"
)

type IOrderUseCase interface {
	PlaceOrder(ctx context.Context, req *dto.PlaceOrderRequest) (*entity.Order, error)
	ListMyOrders(ctx context.Context, req *dto.ListOrdersRequest) ([]*entity.Order, *paging.Pagination, error)
	GetOrderByID(ctx context.Context, id string) (*entity.Order, error)
	UpdateOrder(ctx context.Context, orderID, userID string, status string) (*entity.Order, error)
}

type OrderUseCase struct {
	validator   validation.Validation
	orderRepo   repository.IOrderRepository
	productRepo productRepo.IProductRepository
}

func NewOrderUseCase(
	validator validation.Validation,
	orderRepo repository.IOrderRepository,
	productRepo productRepo.IProductRepository,
) *OrderUseCase {
	return &OrderUseCase{
		validator:   validator,
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (ou *OrderUseCase) PlaceOrder(ctx context.Context, req *dto.PlaceOrderRequest) (*entity.Order, error) {
	if err := ou.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var lines []*entity.OrderLine
	utils.MapStruct(&lines, &req.Lines)

	productMap := make(map[string]*productEntity.Product)
	for _, line := range lines {
		product, err := ou.productRepo.GetProductById(ctx, line.ProductID)
		if err != nil {
			return nil, err
		}
		line.Price = product.Price * float64(line.Quantity)
		productMap[line.ProductID] = product
	}

	order, err := ou.orderRepo.CreateOrder(ctx, req.UserID, lines)
	if err != nil {
		return nil, err
	}

	for _, line := range order.Lines {
		line.Product = productMap[line.ProductID]
	}

	return order, nil
}

func (ou *OrderUseCase) ListMyOrders(ctx context.Context, req *dto.ListOrdersRequest) ([]*entity.Order, *paging.Pagination, error) {
	orders, pagination, err := ou.orderRepo.GetMyOrders(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return orders, pagination, err
}

func (ou *OrderUseCase) GetOrderByID(ctx context.Context, id string) (*entity.Order, error) {
	order, err := ou.orderRepo.GetOrderByID(ctx, id, true)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (ou *OrderUseCase) UpdateOrder(ctx context.Context, orderID, userID string, status string) (*entity.Order, error) {
	order, err := ou.orderRepo.GetOrderByID(ctx, orderID, false)
	if err != nil {
		return nil, err
	}

	if userID != order.UserID {
		return nil, errors.New("permission denied")
	}

	if order.Status == utils.OrderStatusDone || order.Status == utils.OrderStatusCancelled {
		return nil, errors.New("invalid order status")
	}

	statusValue, err := utils.ToOrderStatus(status)
	if err != nil {
		return nil, errors.New("invalid status")
	}

	order.Status = statusValue
	err = ou.orderRepo.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
