package service

import (
	"github.com/atom-apps/door/modules/service/controller"
	"github.com/atom-apps/door/modules/service/dao"
	"github.com/atom-apps/door/modules/service/routes"
	"github.com/atom-apps/door/modules/service/service"

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
