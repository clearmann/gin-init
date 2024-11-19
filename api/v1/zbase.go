package v1

import (
	"errors"
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
	resp := BaseResponse{Code: errorCodeMap[ErrSuccess], Message: ErrSuccess.Error(), Data: data}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = BaseResponse{Code: 0, Message: "", Data: data}
	}
	ctx.JSON(http.StatusOK, resp)
}

// HandleError 不带data数据的错误返回体
func HandleError(ctx *gin.Context, httpCode int, err error) {
	resp := BaseResponse{Code: errorCodeMap[err], Message: err.Error(), Data: nil}
	if _, ok := errorCodeMap[err]; !ok {
		resp = BaseResponse{Code: 500, Message: "unknown error", Data: nil}
	}
	ctx.JSON(httpCode, resp)
}

// HandleErrorWithData 携带data数据的错误返回体
func HandleErrorWithData(ctx *gin.Context, httpCode int, err error, data any) {
	resp := BaseResponse{Code: errorCodeMap[err], Message: err.Error(), Data: data}
	if _, ok := errorCodeMap[err]; !ok {
		resp = BaseResponse{Code: 500, Message: "unknown error", Data: data}
	}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}
