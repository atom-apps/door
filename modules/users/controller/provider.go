package controller

import (
	"github.com/atom-apps/door/modules/users/service"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		permissionRuleSvc *service.PermissionRuleService,
	) (*PermissionRuleController, error) {
		obj := &PermissionRuleController{
			permissionRuleSvc: permissionRuleSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		roleSvc *service.RoleService,
	) (*RoleController, error) {
		obj := &RoleController{
			roleSvc: roleSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		sessionSvc *service.SessionService,
	) (*SessionController, error) {
		obj := &SessionController{
			sessionSvc: sessionSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		tenantSvc *service.TenantService,
	) (*TenantController, error) {
		obj := &TenantController{
			tenantSvc: tenantSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		tokenSvc *service.TokenService,
	) (*TokenController, error) {
		obj := &TokenController{
			tokenSvc: tokenSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		permissionRuleSvc *service.PermissionRuleService,
		userSvc *service.UserService,
	) (*UserController, error) {
		obj := &UserController{
			permissionRuleSvc: permissionRuleSvc,
			userSvc:           userSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
