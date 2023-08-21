// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	 "strings"

	"github.com/atom-apps/door/modules/auth/controller"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeRoutesController(engine fiber.Router, controller *controller.RoutesController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Get(strings.TrimPrefix("/auth/routes", basePath), DataFunc(controller.List))
}