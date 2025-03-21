package repository

import (
	"context"
	"ecommerce_clean/db"
	"ecommerce_clean/internals/order/controller/dto"
	"ecommerce_clean/internals/order/entity"
	"ecommerce_clean/pkgs/paging"
	"ecommerce_clean/utils"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, userID string, lines []*entity.OrderLine) (*entity.Order, error)
	GetOrderByID(ctx context.Context, id string, preload bool) (*entity.Order, error)
	GetMyOrders(ctx context.Context, req *dto.ListOrdersRequest) ([]*entity.Order, *paging.Pagination, error)
	UpdateOrder(ctx context.Context, order *entity.Order) error
}

type OrderRepo struct {
	db db.IDatabase
}

func NewOrderRepository(db db.IDatabase) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, userID string, lines []*entity.OrderLine) (*entity.Order, error) {
	order := new(entity.Order)

	var totalPrice float64
	for _, line := range lines {
		totalPrice += line.Price
	}
	order.TotalPrice = totalPrice
	order.UserID = userID

	handler := func() error {
		return r.createOrder(ctx, order, lines)
	}

	err := r.db.WithTransaction(handler)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepo) createOrder(ctx context.Context, order *entity.Order, lines []*entity.OrderLine) error {
	// Create Order
	if err := r.db.Create(ctx, order); err != nil {
		return err
	}

	// Create order lines
	for _, line := range lines {
		line.OrderID = order.ID
	}
	if err := r.db.CreateInBatches(ctx, &lines, len(lines)); err != nil {
		return err
	}

	utils.MapStruct(&order.Lines, &lines)
	return nil
}

func (r *OrderRepo) GetOrderByID(ctx context.Context, id string, preload bool) (*entity.Order, error) {
	var order entity.Order
	opts := []db.FindOption{
		db.WithQuery(db.NewQuery("id = ?", id)),
	}
	if preload {
		opts = append(opts, db.WithPreload([]string{"Lines", "Lines.Product"}))
	}

	if err := r.db.FindOne(ctx, &order, opts...); err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepo) GetMyOrders(ctx context.Context, req *dto.ListOrdersRequest) ([]*entity.Order, *paging.Pagination, error) {
	query := []db.Query{
		db.NewQuery("user_id = ?", req.UserID),
	}
	if req.Code != "" {
		query = append(query, db.NewQuery("code = ?", req.Code))
	}
	if req.Status != "" {
		query = append(query, db.NewQuery("status = ?", req.Status))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := r.db.Count(ctx, &entity.Order{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	var orders []*entity.Order
	if err := r.db.Find(
		ctx,
		&orders,
		db.WithPreload([]string{"Lines", "Lines.Product"}),
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return orders, pagination, nil
}

func (r *OrderRepo) UpdateOrder(ctx context.Context, order *entity.Order) error {
	return r.db.Update(ctx, order)
}
