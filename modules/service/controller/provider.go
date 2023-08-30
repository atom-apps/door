package controller

import (
	authSvc "github.com/atom-apps/door/modules/auth/service"
	"github.com/atom-apps/door/modules/service/service"
	userSvc "github.com/atom-apps/door/modules/users/service"
	"github.com/atom-providers/captcha"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		captcha *captcha.Captcha,
	) (*CaptchaController, error) {
		obj := &CaptchaController{
			captcha: captcha,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		authSvc *authSvc.AuthService,
		captcha *captcha.Captcha,
		svc *service.SendService,
		userSvc *userSvc.UserService,
	) (*SendController, error) {
		obj := &SendController{
			authSvc: authSvc,
			captcha: captcha,
			svc:     svc,
			userSvc: userSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
