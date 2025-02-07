package v1

import (
	"gin-init/api/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryRequest struct {
	Offset  int      `json:"offset,omitempty"`
	Limit   int      `json:"limit,omitempty"`
	ListAll bool     `json:"list_all,omitempty"`
	OrderBy []string `json:"order_by,omitempty"`
	// OrderType: asc, desc
	OrderType string `json:"order_type,omitempty"`
}
type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// HandleSuccess 成功返回体
func HandleSuccess(ctx *gin.Context, data any) {
	if data == nil {
		data = map[string]any{}
	}

	resp := BaseResponse{Code: errcode.ErrorCodeMap[errcode.Success], Message: errcode.Success.Error(), Data: data}
	if _, ok := errcode.ErrorCodeMap[errcode.Success]; !ok {
		resp = BaseResponse{Code: 0, Message: "", Data: data}
	}
	ctx.JSON(http.StatusOK, resp)
}

// HandleError 不带data数据的错误返回体
func HandleError(ctx *gin.Context, err error) {
	resp := BaseResponse{Code: errcode.ErrorCodeMap[err], Message: err.Error(), Data: nil}
	if _, ok := errcode.ErrorCodeMap[err]; !ok {
		resp = BaseResponse{Code: 500, Message: "unknown error", Data: nil}
	}
	ctx.JSON(http.StatusOK, resp)
}

// HandleErrorWithData 携带data数据的错误返回体
func HandleErrorWithData(ctx *gin.Context, err error, data any) {
	resp := BaseResponse{Code: errcode.ErrorCodeMap[err], Message: err.Error(), Data: data}
	if _, ok := errcode.ErrorCodeMap[err]; !ok {
		resp = BaseResponse{Code: 500, Message: "unknown error", Data: data}
	}
	ctx.JSON(http.StatusOK, resp)
}
