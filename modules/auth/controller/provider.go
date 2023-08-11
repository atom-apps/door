package controller

import (
	"github.com/atom-apps/door/modules/auth/service"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		authSvc *service.AuthService,
		oauth *oauth.Auth,
	) (*AuthController, error) {
		obj := &AuthController{
			authSvc: authSvc,
			oauth:   oauth,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
