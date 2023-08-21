package service

import (
	"github.com/atom-apps/door/modules/user/dao"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-providers/hashids"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/uuid"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		dao *dao.PermissionRuleDao,
		roleDao *dao.RoleDao,
		tenantDao *dao.TenantDao,
	) (*PermissionRuleService, error) {
		obj := &PermissionRuleService{
			dao:       dao,
			roleDao:   roleDao,
			tenantDao: tenantDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		permissionRuleSvc *PermissionRuleService,
		roleDao *dao.RoleDao,
	) (*RoleService, error) {
		obj := &RoleService{
			permissionRuleSvc: permissionRuleSvc,
			roleDao:           roleDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		hash *md5.Hash,
		sessionDao *dao.SessionDao,
		tokenDao *dao.TokenDao,
		userDao *dao.UserDao,
		uuid *uuid.Generator,
	) (*SessionService, error) {
		obj := &SessionService{
			hash:       hash,
			sessionDao: sessionDao,
			tokenDao:   tokenDao,
			userDao:    userDao,
			uuid:       uuid,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		permissionRuleSvc *PermissionRuleService,
		tenantDao *dao.TenantDao,
	) (*TenantService, error) {
		obj := &TenantService{
			permissionRuleSvc: permissionRuleSvc,
			tenantDao:         tenantDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		hash *md5.Hash,
		jwt *jwt.JWT,
		sessionDao *dao.SessionDao,
		tokenDao *dao.TokenDao,
		userDao *dao.UserDao,
		uuid *uuid.Generator,
	) (*TokenService, error) {
		obj := &TokenService{
			hash:       hash,
			jwt:        jwt,
			sessionDao: sessionDao,
			tokenDao:   tokenDao,
			userDao:    userDao,
			uuid:       uuid,
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
		permissionRuleSvc *PermissionRuleService,
		userDao *dao.UserDao,
	) (*UserService, error) {
		obj := &UserService{
			hash:              hash,
			hashID:            hashID,
			permissionRuleSvc: permissionRuleSvc,
			userDao:           userDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
