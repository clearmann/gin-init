package v1

type SendPhoneCodeCategory string

const (
    SendPhoneCodeRegister SendPhoneCodeCategory = "register"
    SendPhoneCodeLogin    SendPhoneCodeCategory = "login"
    SendPhoneCodeVerify   SendPhoneCodeCategory = "verify"
)

type RegisterRequest struct {
    RequestID string `json:"request_id,omitempty"`
    Username  string `json:"username" binding:"required"`
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required"`
}

type LoginRequest struct {
    RequestID string `json:"request_id,omitempty"`
    Name      string `json:"name" binding:"required"`
    Password  string `json:"password" binding:"required"`
}
type LoginResponse struct {
    AccessToken string `json:"access_token"`
}

type UpdateProfileRequest struct {
    RequestID string `json:"request_id,omitempty"`
    Username  string `json:"username,omitempty"`
    Profile   string `json:"profile,omitempty"`
    Email     string `json:"email,omitempty"`
    UUID      uint64 `json:"uuid,omitempty"`
}
type GetProfileRequest struct {
    RequestID string `json:"request_id,omitempty"`
    UUID      uint64 `json:"uuid"`
}
type GetProfileResponse struct {
    Username string `json:"username,omitempty"`
    Profile  string `json:"profile,omitempty"`
    Email    string `json:"email,omitempty"`
    UUID     uint64 `json:"uuid,omitempty"`
    Avatar   string `json:"avatar,omitempty"`
}
type SendPhoneCodeRequest struct {
    Type  string `json:"type"`
    Phone string `json:"phone"`
}

// BindWeChatRequest 绑定第三方微信请求体
type BindWeChatRequest struct {
    Request struct {
        Code string `json:"code"`
    }
}
