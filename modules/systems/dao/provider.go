package dao

import (
	"github.com/atom-apps/door/database/query"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		query *query.Query,
	) (*DictionaryDao, error) {
		obj := &DictionaryDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*MenuDao, error) {
		obj := &MenuDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*RouteDao, error) {
		obj := &RouteDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
