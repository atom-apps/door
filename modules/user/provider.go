package user

import (
	"github.com/atom-apps/door/modules/user/controller"
	"github.com/atom-apps/door/modules/user/dao"
	"github.com/atom-apps/door/modules/user/routes"
	"github.com/atom-apps/door/modules/user/service"

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
