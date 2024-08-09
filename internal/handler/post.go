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

// Create godoc
// @Summary 创建帖子
// @Schemes
// @Description 创建帖子
// @Tags 帖子模块
// @Accept json
// @Produce json
// @Param request body v1.CreatePostRequest true "params"
// @Success 200 {object} v1.BaseResponse
// @Router /register [post]
func (h *PostHandler) Create(ctx *gin.Context) {
    req := new(v1.CreatePostRequest)
    resp := new(v1.BaseResponse)
    if err := ctx.ShouldBindJSON(req); err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }
    if err := h.postService.Create(ctx, req, resp); err != nil {
        h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
        v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// Delete godoc
// @Summary 删除帖子
// @Schemes
// @Description 删除帖子
// @Tags 帖子模块
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *PostHandler) Delete(ctx *gin.Context) {
    req := new(v1.DeletePostRequest)
    resp := new(v1.BaseResponse)
    if err := ctx.ShouldBindJSON(&req); err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    err := h.postService.Delete(ctx, req, resp)
    if err != nil {
        v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// Update godoc
// @Summary 更新帖子信息
// @Schemes
// @Description 更新帖子信息
// @Tags 帖子模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.GetProfileResponse
// @Router /user [get]
func (h *PostHandler) Update(ctx *gin.Context) {
    req := new(v1.UpdatePostRequest)
    resp := new(v1.BaseResponse)
    if err := ctx.ShouldBindJSON(&req); err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    err := h.postService.Update(ctx, req, resp)
    if err != nil {
        v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}
