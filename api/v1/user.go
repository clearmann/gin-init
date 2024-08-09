package v1

import model "gin-init/internal/model/model_type"

type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
    AccessToken string `json:"access_token"`
}

type UpdateProfileRequest struct {
    model.User
}
type GetProfileRequest struct {
    UUID uint64 `json:"uuid"`
}
type GetProfileResponse struct {
    User model.User `json:"user"`
}
