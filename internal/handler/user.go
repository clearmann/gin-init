package handler

import (
	"gin-init/api/errcode"
	"gin-init/api/v1"
	"gin-init/internal/service"
	"gin-init/pkg/utils/validate"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// Register godoc
// @Summary 用户注册
// @Schemes
// @Description 目前只支持邮箱注册
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.RegisterRequest true "params"
// @Success 200 {object} v1.BaseResponse
// @Router /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(v1.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, errcode.ErrBadRequest)
		return
	}
	if !validate.VerifyEmail(req.Email) {
		v1.HandleError(ctx, errcode.ErrEmailFormat)
	}
	if err := h.userService.Register(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		v1.HandleError(ctx, err)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// Login godoc
// @Summary 账号登录
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	req := new(v1.LoginRequest)
	resp := new(v1.LoginResponse)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, errcode.ErrBadRequest)
		return
	}

	err := h.userService.Login(ctx, req, resp)
	if err != nil {
		v1.HandleError(ctx, errcode.ErrNoAuth)
		return
	}
	v1.HandleSuccess(ctx, resp)
}

// GetProfile godoc
// @Summary 获取用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.GetProfileResponse
// @Router /user [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	uuid := GetUUIDFromCtx(ctx)
	if uuid == 0 {
		v1.HandleError(ctx, errcode.ErrNoAuth)
		return
	}
	var req = &v1.GetProfileRequest{
		UUID: uuid,
	}
	resp := new(v1.GetProfileResponse)
	err := h.userService.GetProfile(ctx, req, resp)
	if err != nil {
		v1.HandleError(ctx, errcode.ErrBadRequest)
		return
	}

	v1.HandleSuccess(ctx, resp)
}

// UpdateProfile godoc
// @Summary 修改用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.UpdateProfileRequest true "params"
// @Success 200 {object} v1.BaseResponse
// @Router /user [put]
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	uuid := GetUUIDFromCtx(ctx)
	if uuid == 0 {
		v1.HandleError(ctx, errcode.ErrNoAuth)
		return
	}
	var req *v1.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, errcode.ErrBadRequest)
		return
	}
	req.UUID = uuid
	if err := h.userService.UpdateProfile(ctx, req); err != nil {
		v1.HandleError(ctx, errcode.ErrNoAuth)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// SendPhoneCode godoc
// @Summary 发送验证码
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *UserHandler) SendPhoneCode(ctx *gin.Context) {
	req := new(v1.SendPhoneCodeRequest)
	resp := new(v1.BaseResponse)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, errcode.ErrBadRequest)
		return
	}

	err := h.userService.SendPhoneCode(ctx, req)
	if err != nil {
		v1.HandleError(ctx, errcode.ErrNoAuth)
		return
	}
	v1.HandleSuccess(ctx, resp)
}

// VerifyPhoneCode godoc
// @Summary 验证验证码
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *UserHandler) VerifyPhoneCode(ctx *gin.Context) {
	req := new(v1.LoginRequest)
	resp := new(v1.LoginResponse)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, errcode.ErrBadRequest)
		return
	}

	err := h.userService.Login(ctx, req, resp)
	if err != nil {
		v1.HandleError(ctx, errcode.ErrNoAuth)
		return
	}
	v1.HandleSuccess(ctx, resp)
}

// BindWechat godoc
// @Summary 绑定微信
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.BindWeChatRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *UserHandler) BindWechat(ctx *gin.Context) {
	req := new(v1.BindWeChatRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, errcode.ErrBadRequest)
		return
	}

	err := h.userService.BindWechat(ctx, req)
	if err != nil {
		v1.HandleError(ctx, errcode.ErrNoAuth)
		return
	}
}
