package repository

import (
	"context"
	"ecommerce_clean/configs"
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

func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) ListUsers(ctx context.Context, req *dto.ListUserRequest) ([]*entity.User, *paging.Pagination, error) {
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
	if err := ur.db.Count(ctx, &entity.User{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	var users []*entity.User
	if err := ur.db.Find(
		ctx,
		&users,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (ur *UserRepository) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := ur.db.FindById(ctx, id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	query := db.NewQuery("email = ?", email)
	if err := ur.db.FindOne(ctx, &user, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return ur.db.Create(ctx, user)
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return ur.db.Update(ctx, user)
}

func (ur *UserRepository) DeleteUser(ctx context.Context, user *entity.User) error {
	return ur.db.Delete(ctx, user)
}
