package service

import (
	"github.com/atom-apps/door/modules/tools/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
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

	return nil
}
