// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	 "strings"

	"github.com/atom-apps/door/modules/auth/controller"
	"github.com/atom-apps/door/modules/auth/dto"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeAuthController(engine fiber.Router, controller *controller.AuthController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Post(strings.TrimPrefix("/v1/auth/signup", basePath), DataFunc1(controller.SignUp, Body[dto.SignUpForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/auth/signin", basePath), DataFunc1(controller.SignIn, Body[dto.SignInForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/auth/logout", basePath), Func1(controller.Logout, Body[dto.LogoutForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/auth/refresh-token", basePath), DataFunc1(controller.RefreshToken, Body[dto.RefreshTokenForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/auth/exchange-token-by-code", basePath), DataFunc1(controller.ExchangeTokenByCode, Body[dto.ExchangeTokenByCodeForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/auth/check-reset-password-code", basePath), DataFunc1(controller.CheckResetPasswordCoe, Body[dto.CheckPasswordResetCodeForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/auth/reset-password-by-token", basePath), Func1(controller.ResetPassword, Body[dto.ResetPasswordForm](BodyParamError)))
}
