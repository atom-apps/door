package routes

import (
	"github.com/atom-apps/door/modules/test/controller"
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

func newRoute(svc contracts.HttpService, userController *controller.UserController, roleController *controller.RoleController) contracts.HttpRoute {
	engine := svc.GetEngine().(*fiber.App)
	group := engine.Group("tests")
	log.Infof("register route group: %s", group.(*fiber.Group).Prefix)

	routeRoleController(group, roleController)
	routeUserController(group, userController)
	return nil
}
