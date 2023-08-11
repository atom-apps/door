package errorx

import "errors"

var (
	ErrMissingCodeOrPassword                 = errors.New("缺少验证码或密码")
	ErrorUsernameOrPasswordInvalid           = errors.New("用户名或密码错误")
	ErrorUsernameOrEmailOrPhoneAlreadyExists = errors.New("用户名或邮箱或手机号已存在")
)
