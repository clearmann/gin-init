package handler

import (
    "gin-init/api/v1"
    "gin-init/internal/service"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PostHandler struct {
    *Handler
    postService service.PostService
}

func NewPostHandler(handler *Handler, postService service.PostService) *PostHandler {
    return &PostHandler{
        Handler:     handler,
        postService: postService,
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
// @Success 200 {object} v1.Response
// @Router /register [post]
func (h *PostHandler) Register(ctx *gin.Context) {
    req := new(v1.RegisterRequest)
    if err := ctx.ShouldBindJSON(req); err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }
    if err := h.postService.Register(ctx, req); err != nil {
        h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
        v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
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
func (h *PostHandler) Login(ctx *gin.Context) {
    var req v1.LoginRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    token, err := h.postService.Login(ctx, &req)
    if err != nil {
        v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
        return
    }
    v1.HandleSuccess(ctx, v1.LoginResponseData{
        AccessToken: token,
    })
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
func (h *PostHandler) GetProfile(ctx *gin.Context) {
    uuid := GetUUIDFromCtx(ctx)
    if uuid == 0 {
        v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
        return
    }

    user, err := h.postService.GetProfile(ctx, uuid)
    if err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    v1.HandleSuccess(ctx, user)
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
// @Success 200 {object} v1.Response
// @Router /user [put]
func (h *PostHandler) UpdateProfile(ctx *gin.Context) {
    uuid := GetUUIDFromCtx(ctx)

    var req v1.UpdateProfileRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    if err := h.postService.UpdateProfile(ctx, uuid, &req); err != nil {
        v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
        return
    }

    v1.HandleSuccess(ctx, nil)
}