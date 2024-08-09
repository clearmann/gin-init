package v1

var (
    // common errors
    ErrSuccess             = newError(0, "ok")
    ErrBadRequest          = newError(400, "参数校验错误")
    ErrUnauthorized        = newError(401, "权限校验未通过，请登录后再试")
    ErrNotFound            = newError(404, "未找到，请更换后再试")
    ErrInternalServerError = newError(500, "服务器错误，请稍后再试")

    // more biz errors
    ErrEmailAlreadyUse = newError(1001, "该邮箱已经存在，请更换后再试")
)
