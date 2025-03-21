package usecase

import (
	"context"
	"ecommerce_clean/internals/product/controller/dto"
	"ecommerce_clean/internals/product/entity"
	"ecommerce_clean/internals/product/repository"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/minio"
	"ecommerce_clean/pkgs/paging"
	"ecommerce_clean/pkgs/validation"
	"ecommerce_clean/utils"
)

type IProductUseCase interface {
	ListProducts(ctx context.Context, req *dto.ListProductRequest) ([]*entity.Product, *paging.Pagination, error)
	GetProductById(ctx context.Context, id string) (*entity.Product, error)
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) error
	UpdateProduct(ctx context.Context, id string, req *dto.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, id string) error
}

type ProductUseCase struct {
	validator   validation.Validation
	productRepo repository.IProductRepository
	minioClient *minio.MinioClient
}

func NewProductUseCase(
	validator validation.Validation,
	productRepo repository.IProductRepository,
	minioClient *minio.MinioClient,
) *ProductUseCase {
	return &ProductUseCase{
		validator:   validator,
		productRepo: productRepo,
		minioClient: minioClient,
	}
}

func (pu *ProductUseCase) ListProducts(ctx context.Context, req *dto.ListProductRequest) ([]*entity.Product, *paging.Pagination, error) {
	products, pagination, err := pu.productRepo.ListProducts(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return products, pagination, nil
}

func (pu *ProductUseCase) GetProductById(ctx context.Context, id string) (*entity.Product, error) {
	product, err := pu.productRepo.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) error {
	if err := pu.validator.ValidateStruct(req); err != nil {
		return err
	}

	var imageUrlUpload = ""
	if req.Image != nil {
		avatarURL, err := pu.minioClient.UploadFile(ctx, req.Image, "products")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return err
		}
		imageUrlUpload = avatarURL
	}

	var product entity.Product
	utils.MapStruct(&product, &req)
	product.ImageUrl = imageUrlUpload

	err := pu.productRepo.CreatedProduct(ctx, &product)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return err
	}
	return nil
}

func (pu *ProductUseCase) UpdateProduct(ctx context.Context, id string, req *dto.UpdateProductRequest) error {
	if err := pu.validator.ValidateStruct(req); err != nil {
		return err
	}

	product, err := pu.productRepo.GetProductById(ctx, id)
	if err != nil {
		logger.Errorf("Get fail, error: %s", err)
		return err
	}

	utils.MapStruct(product, req)

	if req.Image != nil {
		avatarURL, err := pu.minioClient.UploadFile(ctx, req.Image, "products")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return err
		}

		pu.minioClient.DeleteFile(ctx, product.ImageUrl)

		product.ImageUrl = avatarURL
	}

	err = pu.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (pu *ProductUseCase) DeleteProduct(ctx context.Context, id string) error {
	product, err := pu.productRepo.GetProductById(ctx, id)
	if err != nil {
		return err
	}

	if err := pu.productRepo.DeleteProduct(ctx, product); err != nil {
		return err
	}

	pu.minioClient.DeleteFile(ctx, product.ImageUrl)

	return nil
}
