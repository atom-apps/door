package routes

import (
	"github.com/atom-apps/door/modules/posts/controller"
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

func newRoute(
	svc contracts.HttpService,
	chapterController *controller.ChapterController,
	bookController *controller.BookController,
	articleController *controller.ArticleController,
) contracts.HttpRoute {
	engine := svc.GetEngine().(*fiber.App)
	group := engine.Group("/v1/posts")
	log.Infof("register route group: %s", group.(*fiber.Group).Prefix)

	routeArticleController(group, articleController)
	routeBookController(group, bookController)
	routeChapterController(group, chapterController)
	return nil
}
