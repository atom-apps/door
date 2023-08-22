package tools

import (
	"github.com/atom-apps/door/modules/tools/controller"
	"github.com/atom-apps/door/modules/tools/dao"
	"github.com/atom-apps/door/modules/tools/routes"
	"github.com/atom-apps/door/modules/tools/service"

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
