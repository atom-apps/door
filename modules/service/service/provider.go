package service

import (
	userSvc "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-providers/app"
	redis "github.com/redis/go-redis/v9"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		app *app.Config,
		cache redis.Cmdable,
		userSvc *userSvc.UserService,
	) (*SendService, error) {
		obj := &SendService{
			app:     app,
			cache:   cache,
			userSvc: userSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
