package service

import (
	"github.com/atom-providers/app"
	redis "github.com/redis/go-redis/v9"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		app *app.Config,
		cache redis.Cmdable,
	) (*SendService, error) {
		obj := &SendService{
			app:   app,
			cache: cache,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
