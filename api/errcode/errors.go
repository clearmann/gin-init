package errcode

import (
	"errors"
)

type Error struct {
	Code    int
	Message string
}

var ErrorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	ErrorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}

var (
	// common errors
	Success          = newError(0, "ok")
	ErrBadRequest    = newError(40000, "请求参数错误")
	ErrNotLogin      = newError(40100, "未登录")
	ErrNoAuth        = newError(40101, "无权限")
	ErrForbidden     = newError(40300, "禁止访问")
	ErrNotFound      = newError(40400, "未找到，请修改后再试")
	ErrInternalError = newError(50000, "系统内部异常")
	ErrOperation     = newError(50001, "操作失败")

	// more biz errors
	ErrEmailAlreadyUse    = newError(10001, "该邮箱已经存在，请修改后再试")
	ErrEmailFormat        = newError(10002, "该邮箱格式错误，请修改后再试")
	ErrUsernameAlreadyUse = newError(10003, "该用户名已经存在，请修改后再试")
	ErrPhoneFormat        = newError(10003, "该手机号格式错误，请修改后再试")
)
