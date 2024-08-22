package v1

var (
    // common errors
    ErrSuccess             = newError(0, "ok")
    ErrBadRequest          = newError(400, "参数校验错误")
    ErrUnauthorized        = newError(401, "权限校验未通过，请登录后再试")
    ErrNotFound            = newError(404, "未找到，请修改后再试")
    ErrInternalServerError = newError(500, "服务器错误，请稍后再试")

    // more biz errors
    ErrEmailAlreadyUse    = newError(1001, "该邮箱已经存在，请修改后再试")
    ErrEmailFormat        = newError(1002, "该邮箱格式错误，请修改后再试")
    ErrUsernameAlreadyUse = newError(1003, "该用户名已经存在，请修改后再试")
    ErrPhoneFormat        = newError(1003, "该手机号格式错误，请修改后再试")
)
