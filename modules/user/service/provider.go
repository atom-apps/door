package service

import (
	"github.com/atom-apps/door/modules/user/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		sessionDao *dao.SessionDao,
	) (*SessionService, error) {
		obj := &SessionService{
			sessionDao: sessionDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		tenantUserDao *dao.TenantUserDao,
	) (*TenantUserService, error) {
		obj := &TenantUserService{
			tenantUserDao: tenantUserDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		tenantDao *dao.TenantDao,
	) (*TenantService, error) {
		obj := &TenantService{
			tenantDao: tenantDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		tokenDao *dao.TokenDao,
	) (*TokenService, error) {
		obj := &TokenService{
			tokenDao: tokenDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		userInfoDao *dao.UserInfoDao,
	) (*UserInfoService, error) {
		obj := &UserInfoService{
			userInfoDao: userInfoDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		userDao *dao.UserDao,
	) (*UserService, error) {
		obj := &UserService{
			userDao: userDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
