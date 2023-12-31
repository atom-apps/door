// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/modules/auth/controller"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routePageController(engine fiber.Router, controller *controller.PageController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Get(strings.TrimPrefix("/auth/signin", groupPrefix), Func1(controller.Signin, String("appName", PathParamError)))
	engine.Get(strings.TrimPrefix("/auth/signup", groupPrefix), Func(controller.Signup))
	engine.Get(strings.TrimPrefix("/auth/reset-password", groupPrefix), Func(controller.ResetPassword))
}
