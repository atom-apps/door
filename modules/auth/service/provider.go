package service

import (
	"github.com/atom-apps/door/modules/user/service"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		userSvc *service.UserService,
	) (*AuthService, error) {
		obj := &AuthService{
			userSvc: userSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
