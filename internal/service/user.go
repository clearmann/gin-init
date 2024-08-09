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
    user, err := s.userRepo.GetByEmail(ctx, req.Email)
    if err != nil {
        return v1.ErrInternalServerError
    }
    if err == nil && user != nil {
        return v1.ErrEmailAlreadyUse
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
    user = &model_type.User{
        UUID:     uuid,
        Email:    req.Email,
        Password: string(hashedPassword),
        Nickname: req.Email,
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
    user, err := s.userRepo.GetByEmail(ctx, req.Email)
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
    resp.User = *user
    return nil
}

func (s *userService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) error {
    user, err := s.userRepo.GetByID(ctx, req.UUID)
    if err != nil {
        return err
    }

    user.Email = req.Email
    user.Nickname = req.Nickname

    if err = s.userRepo.Update(ctx, user); err != nil {
        return err
    }

    return nil
}
