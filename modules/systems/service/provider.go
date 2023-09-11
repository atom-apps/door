package service

import (
	"github.com/atom-apps/door/modules/systems/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		dictionaryDao *dao.DictionaryDao,
	) (*DictionaryService, error) {
		obj := &DictionaryService{
			dictionaryDao: dictionaryDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		locationDao *dao.LocationDao,
	) (*LocationService, error) {
		obj := &LocationService{
			locationDao: locationDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		menuDao *dao.MenuDao,
	) (*MenuService, error) {
		obj := &MenuService{
			menuDao: menuDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

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
