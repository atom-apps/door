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
	engine.Get(strings.TrimPrefix("/auth/signin/:appName", basePath), Func1(controller.SigninPage, String("appName", PathParamError)))
	engine.Get(strings.TrimPrefix("/auth/signup/:appName", basePath), Func1(controller.SignupPage, String("appName", PathParamError)))
	engine.Post(strings.TrimPrefix("/auth/signup", basePath), DataFunc1(controller.SignUp, Body[dto.SignUpForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/auth/signin", basePath), DataFunc1(controller.SignIn, Body[dto.SignInForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/auth/logout", basePath), Func1(controller.Logout, Body[dto.LogoutForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/auth/refresh-token", basePath), DataFunc1(controller.RefreshToken, Body[dto.RefreshTokenForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/auth/exchange-token-by-code", basePath), DataFunc1(controller.ExchangeTokenByCode, Body[dto.ExchangeTokenByCodeForm](BodyParamError)))
}
