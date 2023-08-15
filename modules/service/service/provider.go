package service

import (
	redis "github.com/redis/go-redis/v9"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		cache redis.Cmdable,
	) (*SendService, error) {
		obj := &SendService{
			cache: cache,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
