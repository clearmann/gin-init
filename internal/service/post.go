package service

import (
    "context"
    v1 "gin-init/api/v1"
    "gin-init/internal/repository"
)

type PostService interface {
    Create(ctx context.Context, req *v1.CreatePostRequest, resp *v1.BaseResponse) error
    Delete(ctx context.Context, req *v1.DeletePostRequest, resp *v1.BaseResponse) error
    Update(ctx context.Context, req *v1.UpdatePostRequest, resp *v1.BaseResponse) error
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

func (s *postService) Create(ctx context.Context, req *v1.CreatePostRequest, resp *v1.BaseResponse) error {
    return nil
}

func (s *postService) Delete(ctx context.Context, req *v1.DeletePostRequest, resp *v1.BaseResponse) error {
    return nil
}

func (s *postService) Update(ctx context.Context, req *v1.UpdatePostRequest, resp *v1.BaseResponse) error {
    return nil
}
