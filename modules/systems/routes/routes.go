package routes

import (
	"github.com/atom-apps/door/modules/systems/controller"
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

func newRoute(svc contracts.HttpService, menuController *controller.MenuController, dictionaryController *controller.DictionaryController, routeController *controller.RouteController) contracts.HttpRoute {
	engine := svc.GetEngine().(*fiber.App)
	group := engine.Group("v1/systems")
	log.Infof("register route group: %s", group.(*fiber.Group).Prefix)

	routeRouteController(group, routeController)
	routeDictionaryController(group, dictionaryController)
	routeMenuController(group, menuController)
	return nil
}
