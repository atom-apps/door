package service

import (
	serviceSvc "github.com/atom-apps/door/modules/service/service"
	"github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-providers/uuid"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		hash *bcrypt.Hash,
		sendSvc *serviceSvc.SendService,
		userSvc *service.UserService,
		uuid *uuid.Generator,
	) (*AuthService, error) {
		obj := &AuthService{
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
