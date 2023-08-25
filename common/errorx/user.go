package errorx

import (
	"net/http"

	"github.com/rogeecn/fen"
)

var (
	ErrUserNotExists                         = fen.NewBusError(http.StatusNotFound, http.StatusNotFound, "用户不存在")
	ErrMissingCodeOrPassword                 = fen.NewBusError(http.StatusBadRequest, http.StatusBadGateway, "缺少验证码或密码")
	ErrorUsernameOrPasswordInvalid           = fen.NewBusError(http.StatusBadRequest, http.StatusBadGateway, "用户名或密码错误")
	ErrorUsernameOrEmailOrPhoneAlreadyExists = fen.NewBusError(http.StatusBadRequest, http.StatusBadGateway, "用户名或邮箱或手机号已存在")
)
