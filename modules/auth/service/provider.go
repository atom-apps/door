package service

import (
	serviceSvc "github.com/atom-apps/door/modules/service/service"
	"github.com/atom-apps/door/modules/users/service"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/uuid"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		auth *oauth.Auth,
		hash *bcrypt.Hash,
		sendSvc *serviceSvc.SendService,
		userSvc *service.UserService,
		uuid *uuid.Generator,
	) (*AuthService, error) {
		obj := &AuthService{
			auth:    auth,
			hash:    hash,
			sendSvc: sendSvc,
			userSvc: userSvc,
			uuid:    uuid,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
