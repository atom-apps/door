package controller

import (
	"github.com/atom-apps/door/modules/tools/service"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		locationSvc *service.LocationService,
	) (*LocationController, error) {
		obj := &LocationController{
			locationSvc: locationSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
