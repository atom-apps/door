package controller

import (
	"github.com/atom-apps/door/modules/systems/service"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		routeSvc *service.RouteService,
	) (*RouteController, error) {
		obj := &RouteController{
			routeSvc: routeSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
