package oauth

import "errors"

var (
	ErrInvalidApp = errors.New("找不到应用")

	ErrEmailRequired = errors.New("电子邮箱不能为空")
	ErrEmailInvalid  = errors.New("电子邮箱不合法")

	ErrPhoneRequired = errors.New("手机号码不能为空")
	ErrPhoneInvalid  = errors.New("手机号码不合法")

	ErrUsernameRequired = errors.New("用户名不能为空")
	ErrUsernameInvalid  = errors.New("用户名不合法")

	ErrPasswordRequired = errors.New("密码不能为空")
	ErrPasswordInvalid  = errors.New("密码不合法")

	ErrVerifyCodeInvalid = errors.New("验证码不正确")

	ErrPhoneExists = errors.New("手机号已存在")
	ErrEmailExists = errors.New("邮箱已存在")
)
