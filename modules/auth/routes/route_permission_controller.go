// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/modules/auth/controller"
	"github.com/atom-apps/door/modules/auth/dto"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routePermissionController(engine fiber.Router, controller *controller.PermissionController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Post(strings.TrimPrefix("/v1/auth/permission/check", groupPrefix), Func1(controller.Check, Body[dto.PermissionCheckForm](BodyParamError)))
}
