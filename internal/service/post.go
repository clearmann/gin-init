package service

import (
    "context"
    v1 "gin-init/api/v1"
    "gin-init/internal/repository"
)

type PostService interface {
    Register(ctx context.Context, req *v1.RegisterRequest) error
    Login(ctx context.Context, req *v1.LoginRequest) (string, error)
    GetProfile(ctx context.Context, uuid uint64) (*v1.GetProfileResponseData, error)
    UpdateProfile(ctx context.Context, uuid uint64, req *v1.UpdateProfileRequest) error
}

func NewPostService(
    service *Service,
    postRepo repository.PostRepository,
) PostService {
    return &postService{
        postRepo: postRepo,
        Service:  service,
    }
}

type postService struct {
    postRepo repository.PostRepository
    *Service
}

func (s *postService) Register(ctx context.Context, req *v1.RegisterRequest) error {
    return nil
}

func (s *postService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
    return "", nil
}

func (s *postService) GetProfile(ctx context.Context, uuid uint64) (*v1.GetProfileResponseData, error) {
    return nil, nil
}

func (s *postService) UpdateProfile(ctx context.Context, uuid uint64, req *v1.UpdateProfileRequest) error {
    return nil
}
