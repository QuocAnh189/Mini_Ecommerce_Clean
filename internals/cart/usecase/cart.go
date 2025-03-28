package usecase

import (
	"context"
	"ecommerce_clean/utils"

	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/validation"

	"ecommerce_clean/internals/cart/controller/dto"
	"ecommerce_clean/internals/cart/entity"
	"ecommerce_clean/internals/cart/repository"
)

type ICartUseCase interface {
	GetCartByUserID(ctx context.Context, userID string) (*entity.Cart, error)
	AddProduct(ctx context.Context, req *dto.AddProductRequest) error
	RemoveProduct(ctx context.Context, req *dto.RemoveProductRequest) error
}

type CartUseCase struct {
	validator validation.Validation
	cartRepo  repository.ICartRepository
}

func NewCartUseCase(
	validator validation.Validation,
	cartRepo repository.ICartRepository,
) *CartUseCase {
	return &CartUseCase{
		validator: validator,
		cartRepo:  cartRepo,
	}
}

func (cu *CartUseCase) GetCartByUserID(ctx context.Context, userID string) (*entity.Cart, error) {
	cart, err := cu.cartRepo.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (cu *CartUseCase) AddProduct(ctx context.Context, req *dto.AddProductRequest) error {
	if err := cu.validator.ValidateStruct(req); err != nil {
		return err
	}

	var cartLine entity.CartLine
	utils.MapStruct(&cartLine, &req)

	err := cu.cartRepo.CreateCartLine(ctx, &cartLine)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return err
	}
	return nil
}

func (cu *CartUseCase) RemoveProduct(ctx context.Context, req *dto.RemoveProductRequest) error {
	cartLine, err := cu.cartRepo.GetCartLineByProductIDAndCartID(ctx, req)
	if err != nil {
		return err
	}

	if err := cu.cartRepo.RemoveCartLine(ctx, cartLine); err != nil {
		return err
	}

	return nil
}
