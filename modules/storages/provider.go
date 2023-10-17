package storages

import (
	"github.com/atom-apps/door/modules/storages/controller"
	"github.com/atom-apps/door/modules/storages/dao"
	"github.com/atom-apps/door/modules/storages/routes"
	"github.com/atom-apps/door/modules/storages/service"

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
