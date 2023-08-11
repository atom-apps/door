package errorx

import "errors"

var (
	ErrMissingCodeOrPassword       = errors.New("缺少验证码或密码")
	ErrorUsernameOrPasswordInvalid = errors.New("用户名或密码错误")
)
