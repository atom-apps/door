package service

import (
	"github.com/atom-apps/door/modules/user/dao"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-apps/door/providers/jwt"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-providers/hashids"
	"github.com/atom-providers/uuid"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		dao *dao.PermissionRuleDao,
	) (*PermissionRuleService, error) {
		obj := &PermissionRuleService{
			dao: dao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		roleUserDao *dao.RoleUserDao,
	) (*RoleUserService, error) {
		obj := &RoleUserService{
			roleUserDao: roleUserDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		roleDao *dao.RoleDao,
		roleUserDao *dao.RoleUserDao,
	) (*RoleService, error) {
		obj := &RoleService{
			roleDao:     roleDao,
			roleUserDao: roleUserDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		hash *md5.Hash,
		sessionDao *dao.SessionDao,
		tokenDao *dao.TokenDao,
		uuid *uuid.Generator,
	) (*SessionService, error) {
		obj := &SessionService{
			hash:       hash,
			sessionDao: sessionDao,
			tokenDao:   tokenDao,
			uuid:       uuid,
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
		hash *md5.Hash,
		jwt *jwt.JWT,
		tokenDao *dao.TokenDao,
		userDao *dao.UserDao,
		uuid *uuid.Generator,
	) (*TokenService, error) {
		obj := &TokenService{
			hash:     hash,
			jwt:      jwt,
			tokenDao: tokenDao,
			userDao:  userDao,
			uuid:     uuid,
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
		hash *bcrypt.Hash,
		hashID *hashids.HashID,
		userDao *dao.UserDao,
	) (*UserService, error) {
		obj := &UserService{
			hash:    hash,
			hashID:  hashID,
			userDao: userDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
