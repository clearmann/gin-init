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

// List godoc
// @Summary 列出帖子信息
// @Schemes
// @Description 列出帖子信息
// @Tags 帖子模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.ListPostRequest true "params"
// @Success 200 {object} v1.ListPostResponse
// @Router /post/list [post]
func (h *PostHandler) List(ctx *gin.Context) {
    req := new(v1.ListPostRequest)
    resp := new(v1.ListPostResponse)
    if err := ctx.ShouldBindJSON(&req); err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    err := h.postService.List(ctx, req, resp)
    if err != nil {
        v1.HandleError(ctx, http.StatusOK, err, nil)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}
