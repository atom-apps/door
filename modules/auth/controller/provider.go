package controller

import (
	"github.com/atom-apps/door/modules/auth/service"
	serviceSvc "github.com/atom-apps/door/modules/service/service"
	userSvc "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/uuid"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		authSvc *service.AuthService,
		oauth *oauth.Auth,
		permissionSvc *userSvc.PermissionRuleService,
		sendSvc *serviceSvc.SendService,
		sessionSvc *userSvc.SessionService,
		tenantSvc *userSvc.TenantService,
		tokenSvc *userSvc.TokenService,
		userSvc *userSvc.UserService,
	) (*AuthController, error) {
		obj := &AuthController{
			authSvc:       authSvc,
			oauth:         oauth,
			permissionSvc: permissionSvc,
			sendSvc:       sendSvc,
			sessionSvc:    sessionSvc,
			tenantSvc:     tenantSvc,
			tokenSvc:      tokenSvc,
			userSvc:       userSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		casbin *casbin.Casbin,
		oauth *oauth.Auth,
		sessionSvc *userSvc.SessionService,
		tokenSvc *userSvc.TokenService,
		userSvc *userSvc.UserService,
		uuid *uuid.Generator,
	) (*PageController, error) {
		obj := &PageController{
			casbin:     casbin,
			oauth:      oauth,
			sessionSvc: sessionSvc,
			tokenSvc:   tokenSvc,
			userSvc:    userSvc,
			uuid:       uuid,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func() (*RoutesController, error) {
		obj := &RoutesController{}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
