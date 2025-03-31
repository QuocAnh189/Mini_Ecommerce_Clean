package repository

import (
	"context"
	"ecommerce_clean/db"
	"ecommerce_clean/internals/cart/entity"
)

type ICartRepository interface {
	GetCartByUserID(ctx context.Context, userID string) (*entity.Cart, error)
	GetCartLineByProductIDAndCartID(ctx context.Context, cartID string, productID string) (*entity.CartLine, error)
	CreateCartLine(ctx context.Context, cartLine *entity.CartLine) error
	UpdateCartLine(ctx context.Context, cartLine *entity.CartLine) error
	RemoveCartLine(ctx context.Context, cartLine *entity.CartLine) error
}

type CartRepository struct {
	db db.IDatabase
}

func NewCartRepository(db db.IDatabase) *CartRepository {
	return &CartRepository{db: db}
}

func (cr *CartRepository) GetCartByUserID(ctx context.Context, userID string) (*entity.Cart, error) {
	var cart entity.Cart
	opts := []db.FindOption{
		db.WithQuery(db.NewQuery("user_id = ?", userID)),
	}
	opts = append(opts, db.WithPreload([]string{"User", "Lines.Product"}))

	if err := cr.db.FindOne(ctx, &cart, opts...); err != nil {
		return nil, err
	}

	return &cart, nil
}

func (cr *CartRepository) GetCartLineByProductIDAndCartID(ctx context.Context, cartID string, productID string) (*entity.CartLine, error) {
	var cartLine entity.CartLine
	opts := []db.FindOption{
		db.WithQuery(db.NewQuery("cart_id = ?", cartID)),
		db.WithQuery(db.NewQuery("product_id = ?", productID)),
	}

	if err := cr.db.FindOne(ctx, &cartLine, opts...); err != nil {
		return nil, err
	}

	return &cartLine, nil
}

func (cr *CartRepository) CreateCartLine(ctx context.Context, cartLine *entity.CartLine) error {
	return cr.db.Create(ctx, cartLine)
}

func (cr *CartRepository) UpdateCartLine(ctx context.Context, cartLine *entity.CartLine) error {
	return cr.db.Update(ctx, cartLine)
}

func (cr *CartRepository) RemoveCartLine(ctx context.Context, cartLine *entity.CartLine) error {
	return cr.db.Delete(ctx, cartLine)
}
