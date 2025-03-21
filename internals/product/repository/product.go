package repository

import (
	"context"
	"ecommerce_clean/configs"
	"ecommerce_clean/db"
	"ecommerce_clean/internals/product/controller/dto"
	"ecommerce_clean/internals/product/entity"
	"ecommerce_clean/pkgs/paging"
)

type IProductRepository interface {
	ListProducts(ctx context.Context, req *dto.ListProductRequest) ([]*entity.Product, *paging.Pagination, error)
	GetProductById(ctx context.Context, id string) (*entity.Product, error)
	CreatedProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, product *entity.Product) error
}

type ProductRepository struct {
	db db.IDatabase
}

func NewProductRepository(db db.IDatabase) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) ListProducts(ctx context.Context, req *dto.ListProductRequest) ([]*entity.Product, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]db.Query, 0)

	if req.Search != "" {
		query = append(query, db.NewQuery("name ILIKE ?", "%"+req.Search+"%"))
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := pr.db.Count(ctx, &entity.Product{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	var products []*entity.Product
	if err := pr.db.Find(
		ctx,
		&products,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (pr *ProductRepository) GetProductById(ctx context.Context, id string) (*entity.Product, error) {
	var product entity.Product
	if err := pr.db.FindById(ctx, id, &product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) CreatedProduct(ctx context.Context, product *entity.Product) error {
	return pr.db.Create(ctx, product)
}

func (pr *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {
	return pr.db.Update(ctx, product)
}

func (pr *ProductRepository) DeleteProduct(ctx context.Context, product *entity.Product) error {
	return pr.db.Delete(ctx, product)
}
