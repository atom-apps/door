package systems

import (
	"github.com/atom-apps/door/modules/systems/controller"
	"github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/systems/routes"
	"github.com/atom-apps/door/modules/systems/service"

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
