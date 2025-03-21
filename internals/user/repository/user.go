package repository

import (
	"context"
	"ecommerce_clean/db"
	"ecommerce_clean/internals/user/controller/dto"
	"ecommerce_clean/internals/user/entity"
	"ecommerce_clean/pkgs/paging"
)

type IUserRepository interface {
	ListUsers(ctx context.Context, req *dto.ListUserRequest) ([]*entity.User, *paging.Pagination, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user *entity.User) error
}

type UserRepository struct {
	db db.IDatabase
}

func NewCardRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) ListUsers(ctx context.Context, req *dto.ListUserRequest) ([]*entity.User, *paging.Pagination, error) {
	return nil, nil, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.FindById(ctx, id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	query := db.NewQuery("email = ?", email)
	if err := r.db.FindOne(ctx, &user, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.Create(ctx, user)
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	return r.db.Update(ctx, user)
}

func (r *UserRepository) Delete(ctx context.Context, user *entity.User) error {
	return r.db.Delete(ctx, user)
}
