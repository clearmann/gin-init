package repository

import (
    "context"
    "errors"
    v1 "gin-init/api/v1"
    "gin-init/internal/model/model_type"

    "gorm.io/gorm"
)

type UserRepository interface {
    Create(ctx context.Context, user *model_type.User) error
    Update(ctx context.Context, user *model_type.User) error
    GetByID(ctx context.Context, uuid uint64) (*model_type.User, error)
    GetByEmail(ctx context.Context, email string) (*model_type.User, error)
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

func (r *userRepository) Update(ctx context.Context, user *model_type.User) error {
    if err := r.DB(ctx).Save(user).Error; err != nil {
        return err
    }
    return nil
}

func (r *userRepository) GetByID(ctx context.Context, uuid uint64) (*model_type.User, error) {
    var user model_type.User
    if err := r.DB(ctx).Where("uuid = ?", uuid).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, v1.ErrNotFound
        }
        return nil, err
    }
    return &user, nil
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
