package v1

type RegisterRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
    Name     string `json:"name" binding:"required"`
    Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
    AccessToken string `json:"access_token"`
}

type UpdateProfileRequest struct {
    Username string `json:"username,omitempty"`
    Profile  string `json:"profile,omitempty"`
    Email    string `json:"email,omitempty"`
    UUID     uint64 `json:"uuid,omitempty"`
}
type GetProfileRequest struct {
    UUID uint64 `json:"uuid"`
}
type GetProfileResponse struct {
    Username string `json:"username,omitempty"`
    Profile  string `json:"profile,omitempty"`
    Email    string `json:"email,omitempty"`
    UUID     uint64 `json:"uuid,omitempty"`
    Avatar   string `json:"avatar,omitempty"`
}
