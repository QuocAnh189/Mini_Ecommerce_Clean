package usecase

import (
	"context"
	"ecommerce_clean/internals/user/controller/dto"
	"ecommerce_clean/internals/user/entity"
	"ecommerce_clean/internals/user/repository"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/minio"
	"ecommerce_clean/pkgs/paging"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/token"
	"ecommerce_clean/pkgs/validation"
	"ecommerce_clean/utils"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, *entity.User, error)
	SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, *entity.User, error)
	SignOut(ctx context.Context, userID string, jit string) error
	ListUsers(ctx context.Context, req *dto.ListUserRequest) ([]*entity.User, *paging.Pagination, error)
	GetUserById(ctx context.Context, userID string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type UserUseCase struct {
	validator   validation.Validation
	userRepo    repository.IUserRepository
	minioClient *minio.MinioClient
	cache       redis.IRedis
	token       token.IMarker
}

func NewUserUseCase(
	validator validation.Validation,
	userRepo repository.IUserRepository,
	minioClient *minio.MinioClient,
	cache redis.IRedis,
	token token.IMarker,
) *UserUseCase {
	return &UserUseCase{
		validator:   validator,
		userRepo:    userRepo,
		minioClient: minioClient,
		cache:       cache,
		token:       token,
	}
}

func (u *UserUseCase) SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, *entity.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return "", "", nil, err
	}
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
		return "", "", nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", nil, errors.New("wrong password")
	}

	tokenData := token.AuthPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	accessToken := u.token.GenerateAccessToken(&tokenData)
	refreshToken := u.token.GenerateRefreshToken(&tokenData)

	return accessToken, refreshToken, user, nil
}

func (u *UserUseCase) SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, *entity.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return "", "", nil, err
	}

	var avatarUrlUpload = ""
	if req.Avatar != nil {
		avatarURL, err := u.minioClient.UploadFile(ctx, req.Avatar, "users")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return "", "", nil, err
		}
		avatarUrlUpload = avatarURL
	}

	var user *entity.User
	utils.MapStruct(&user, &req)
	user.AvatarUrl = avatarUrlUpload

	err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return "", "", nil, err
	}

	tokenData := token.AuthPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	accessToken := u.token.GenerateAccessToken(&tokenData)
	refreshToken := u.token.GenerateRefreshToken(&tokenData)

	return accessToken, refreshToken, user, nil
}

func (u *UserUseCase) SignOut(ctx context.Context, userID string, jit string) error {
	value := `{"status": "blacklisted"}`

	// err := u.cache.Set(fmt.Sprintf("blacklist:%s", strings.ReplaceAll(token, " ", "_")), value)
	err := u.cache.Set(fmt.Sprintf("blacklist:%s_%s", userID, jit), value)
	if err != nil {
		logger.Error("Failed to blacklist token: ", err)
		return err
	}

	logger.Info("User signed out successfully")
	return nil
}

func (u *UserUseCase) ListUsers(ctx context.Context, req *dto.ListUserRequest) ([]*entity.User, *paging.Pagination, error) {
	users, pagination, err := u.userRepo.ListUsers(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return users, pagination, nil
}

func (u *UserUseCase) GetUserById(ctx context.Context, userID string) (*entity.User, error) {
	card, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	if err := u.userRepo.DeleteUser(ctx, user); err != nil {
		return err
	}

	u.minioClient.DeleteFile(ctx, user.AvatarUrl)

	return nil
}
