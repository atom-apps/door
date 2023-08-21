// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	 "strings"

	"github.com/atom-apps/door/modules/user/controller"
	"github.com/atom-apps/door/common"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routePermissionRuleController(engine fiber.Router, controller *controller.PermissionRuleController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Put(strings.TrimPrefix("/v1/permissions/attach/:roleId<int>/:tenantId", basePath), Func3(controller.AttachUsers, Integer[int64]("id", PathParamError), Integer[int64]("tenantID", PathParamError), Body[common.IDsForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/v1/permissions/detach/:roleId<int>/:tenantId", basePath), Func3(controller.DetachUsers, Integer[int64]("id", PathParamError), Integer[int64]("tenantID", PathParamError), Body[common.IDsForm](BodyParamError)))
}
