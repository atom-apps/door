package routes

import (
	"github.com/atom-apps/door/modules/users/controller"
	"github.com/atom-providers/log"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	return container.Container.Provide(newRoute, atom.GroupRoutes)
}

func newRoute(svc contracts.HttpService, roleController *controller.RoleController, permissionRuleController *controller.PermissionRuleController, userController *controller.UserController, tokenController *controller.TokenController, tenantController *controller.TenantController, sessionController *controller.SessionController) contracts.HttpRoute {
	engine := svc.GetEngine().(*fiber.App)
	group := engine.Group("v1")
	log.Infof("register route group: %s", group.(*fiber.Group).Prefix)

	routeSessionController(group, sessionController)
	routeTenantController(group, tenantController)
	routeTokenController(group, tokenController)
	routeUserController(group, userController)
	routeRoleController(group, roleController)
	routePermissionRuleController(group, permissionRuleController)
	return nil
}
