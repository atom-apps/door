package posts

import (
	"github.com/atom-apps/door/modules/posts/controller"
	"github.com/atom-apps/door/modules/posts/dao"
	"github.com/atom-apps/door/modules/posts/routes"
	"github.com/atom-apps/door/modules/posts/service"

	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: dao.Provide},
		{Provider: service.Provide},
		{Provider: controller.Provide},
		{Provider: routes.Provide},
	}
}
