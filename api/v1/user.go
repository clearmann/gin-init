package v1

type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
    Password string `json:"password" binding:"required" example:"123456"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
    Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
    AccessToken string `json:"access_token"`
}
type LoginResponse struct {
    Response
    Data LoginResponseData
}

type UpdateProfileRequest struct {
    Nickname string `json:"nickname" example:"alan"`
    Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}
type GetProfileResponseData struct {
    UUID     uint64 `json:"uuid"`
    Nickname string `json:"nickname" example:"alan"`
}
type GetProfileResponse struct {
    Response
    Data GetProfileResponseData
}
