package routes

import (
	"github.com/atom-apps/door/modules/auth/controller"
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

func newRoute(svc contracts.HttpService, permissionController *controller.PermissionController, routesController *controller.RoutesController, pageController *controller.PageController, authController *controller.AuthController) contracts.HttpRoute {
	engine := svc.GetEngine().(*fiber.App)
	group := engine.Group("v1")
	log.Infof("register route group: %s", group.(*fiber.Group).Prefix)

	routePageController(engine.Group("auth"), pageController)
	routeAuthController(group, authController)
	routePageController(group, pageController)
	routeRoutesController(group, routesController)
	routePermissionController(group, permissionController)
	return nil
}
