package service

import (
    "context"
    v1 "gin-init/api/v1"
    "gin-init/internal/model/model_type"
    "gin-init/internal/repository"
    "log"
    "time"

    "golang.org/x/crypto/bcrypt"
)

type UserService interface {
    Register(ctx context.Context, req *v1.RegisterRequest) error
    Login(ctx context.Context, req *v1.LoginRequest, resp *v1.LoginResponse) error
    GetProfile(ctx context.Context, req *v1.GetProfileRequest, resp *v1.GetProfileResponse) error
    UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) error
}

func NewUserService(
    service *Service,
    userRepo repository.UserRepository,
) UserService {
    return &userService{
        userRepo: userRepo,
        Service:  service,
    }
}

type userService struct {
    userRepo repository.UserRepository
    *Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
    // check username
    isExist, err := s.userRepo.ExistUserByEmail(ctx, req.Email)
    if err != nil {
        return v1.ErrInternalServerError
    }
    if !isExist {
        return v1.ErrEmailAlreadyUse
    }

    isExist, err = s.userRepo.ExistUserByUsername(ctx, req.Email)
    if err != nil {
        return v1.ErrInternalServerError
    }
    if !isExist {
        return v1.ErrUsernameAlreadyUse
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Generate password error:", err)
        return err
    }
    // Generate uuid
    uuid, err := s.sid.GenUint64()
    if err != nil {
        log.Println("Generate uuid error:", err)
        return err
    }
    var user = &model_type.User{
        UUID:     uuid,
        Email:    req.Email,
        Password: string(hashedPassword),
        Username: req.Username,
    }
    // Transaction demo
    err = s.tm.Transaction(ctx, func(ctx context.Context) error {
        // Create a user
        if err = s.userRepo.Create(ctx, user); err != nil {
            log.Println("Create user error:", err)
            return err
        }
        return nil
    })
    return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest, resp *v1.LoginResponse) error {
    user, err := s.userRepo.GetByEmailORUsername(ctx, req.Name)
    if err != nil || user == nil {
        return v1.ErrUnauthorized
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        s.logger.Info("")
        return err
    }
    token, err := s.jwt.GenToken(user.UUID, time.Now().Add(time.Hour*24*90))
    if err != nil {
        return err
    }
    resp.AccessToken = token
    return nil
}

func (s *userService) GetProfile(ctx context.Context, req *v1.GetProfileRequest, resp *v1.GetProfileResponse) error {
    user, err := s.userRepo.GetByID(ctx, req.UUID)
    if err != nil {
        return err
    }
    resp.Avatar = user.Avatar
    resp.Email = user.Email
    resp.Profile = user.Profile
    resp.UUID = user.UUID
    resp.Username = user.Username
    return nil
}

func (s *userService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) error {
    isExist, err := s.userRepo.ExistUserByUUID(ctx, req.UUID)
    if err != nil {
        return err
    }
    if !isExist {
        return v1.ErrNotFound
    }
    if err = s.userRepo.Update(ctx, req); err != nil {
        return err
    }

    return nil
}
