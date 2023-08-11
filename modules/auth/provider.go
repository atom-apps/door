package auth

import (
	"github.com/atom-apps/door/modules/auth/controller"
	"github.com/atom-apps/door/modules/auth/dao"
	"github.com/atom-apps/door/modules/auth/routes"
	"github.com/atom-apps/door/modules/auth/service"

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
