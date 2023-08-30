package test

import (
	"github.com/atom-apps/door/modules/test/controller"
	"github.com/atom-apps/door/modules/test/dao"
	"github.com/atom-apps/door/modules/test/routes"
	"github.com/atom-apps/door/modules/test/service"

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
