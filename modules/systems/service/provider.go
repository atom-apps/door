package service

import (
	"github.com/atom-apps/door/modules/systems/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		routeDao *dao.RouteDao,
	) (*RouteService, error) {
		obj := &RouteService{
			routeDao: routeDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
