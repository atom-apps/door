package user

import (
	"github.com/atom-apps/door/modules/users/controller"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/routes"
	"github.com/atom-apps/door/modules/users/service"

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
