package controller

import (
	"github.com/atom-apps/door/modules/auth/service"
	serviceSvc "github.com/atom-apps/door/modules/service/service"
	systemSvc "github.com/atom-apps/door/modules/systems/service"
	userSvc "github.com/atom-apps/door/modules/users/service"
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
		sendSvc *serviceSvc.SendService,
		sessionSvc *userSvc.SessionService,
		tenantSvc *userSvc.TenantService,
		tokenSvc *userSvc.TokenService,
		userSvc *userSvc.UserService,
		userTenantRoleSvc *userSvc.UserTenantRoleService,
	) (*AuthController, error) {
		obj := &AuthController{
			authSvc:           authSvc,
			oauth:             oauth,
			sendSvc:           sendSvc,
			sessionSvc:        sessionSvc,
			tenantSvc:         tenantSvc,
			tokenSvc:          tokenSvc,
			userSvc:           userSvc,
			userTenantRoleSvc: userTenantRoleSvc,
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

	if err := container.Container.Provide(func(
		casbin *casbin.Casbin,
	) (*PermissionController, error) {
		obj := &PermissionController{
			casbin: casbin,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		casbinSvc *userSvc.CasbinService,
		permissionSvc *userSvc.PermissionService,
		routeSvc *systemSvc.RouteService,
	) (*RoutesController, error) {
		obj := &RoutesController{
			casbinSvc:     casbinSvc,
			permissionSvc: permissionSvc,
			routeSvc:      routeSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
