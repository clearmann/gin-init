package repository

import (
	"context"
	"errors"
	"gin-init/api/errcode"
	"gin-init/api/v1"
	"gin-init/internal/model/model_type"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model_type.User) error
	Update(ctx context.Context, req *v1.UpdateProfileRequest) error
	GetByID(ctx context.Context, uuid uint64) (*model_type.User, error)
	GetByEmail(ctx context.Context, email string) (*model_type.User, error)
	GetByEmailORUsername(ctx context.Context, name string) (*model_type.User, error)
	ExistUserByUUID(ctx context.Context, uuid uint64) (bool, error)
	ExistUserByEmail(ctx context.Context, email string) (bool, error)
	ExistUserByUsername(ctx context.Context, username string) (bool, error)
	Set(ctx context.Context, key, code string, t time.Duration) error
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{Repository: r}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) Create(ctx context.Context, user *model_type.User) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, req *v1.UpdateProfileRequest) error {
	if err := r.DB(ctx).
		Where("uuid = ?", req.UUID).
		Updates(&model_type.User{Username: req.Username, Email: req.Email, Profile: req.Profile}).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, uuid uint64) (*model_type.User, error) {
	var user model_type.User
	if err := r.DB(ctx).Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) ExistUserByEmail(ctx context.Context, email string) (bool, error) {
	var total *int64
	if err := r.DB(ctx).Where("email = ?", email).Count(total).Error; err != nil {
		return false, err
	}
	return *total > 0, nil
}
func (r *userRepository) ExistUserByUsername(ctx context.Context, username string) (bool, error) {
	var total *int64
	if err := r.DB(ctx).Where("username = ?", username).Count(total).Error; err != nil {
		return false, err
	}
	return *total > 0, nil
}
func (r *userRepository) ExistUserByUUID(ctx context.Context, uuid uint64) (bool, error) {
	var total *int64
	if err := r.DB(ctx).Where("uuid = ?", uuid).Count(total).Error; err != nil {
		return false, err
	}
	return *total > 0, nil
}
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model_type.User, error) {
	var user model_type.User
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetByEmailORUsername(ctx context.Context, name string) (*model_type.User, error) {
	var user model_type.User
	if err := r.DB(ctx).Where("email = ? or username", name, name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) Set(ctx context.Context, key, code string, t time.Duration) error {
	r.rdb.Set(ctx, key, code, t)
	return nil
}
