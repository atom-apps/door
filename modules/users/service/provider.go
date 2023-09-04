package service

import (
	systemDao "github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-apps/door/providers/md5"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/hashids"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/uuid"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		permissionDao *dao.PermissionDao,
		roleDao *dao.RoleDao,
		routeDao *systemDao.RouteDao,
		tenantDao *dao.TenantDao,
	) (*PermissionService, error) {
		obj := &PermissionService{
			permissionDao: permissionDao,
			roleDao:       roleDao,
			routeDao:      routeDao,
			tenantDao:     tenantDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		permissionSvc *PermissionService,
		roleDao *dao.RoleDao,
		userTenantRoleSvc *UserTenantRoleService,
	) (*RoleService, error) {
		obj := &RoleService{
			permissionSvc:     permissionSvc,
			roleDao:           roleDao,
			userTenantRoleSvc: userTenantRoleSvc,
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
		permissionSvc *PermissionService,
		tenantDao *dao.TenantDao,
		userTenantRoleSvc *UserTenantRoleService,
	) (*TenantService, error) {
		obj := &TenantService{
			permissionSvc:     permissionSvc,
			tenantDao:         tenantDao,
			userTenantRoleSvc: userTenantRoleSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		UserTenantRoleSvc *UserTenantRoleService,
		auth *oauth.Auth,
		hash *md5.Hash,
		jwt *jwt.JWT,
		roleDao *dao.RoleDao,
		sessionDao *dao.SessionDao,
		tokenDao *dao.TokenDao,
		userDao *dao.UserDao,
		uuid *uuid.Generator,
	) (*TokenService, error) {
		obj := &TokenService{
			UserTenantRoleSvc: UserTenantRoleSvc,
			auth:              auth,
			hash:              hash,
			jwt:               jwt,
			roleDao:           roleDao,
			sessionDao:        sessionDao,
			tokenDao:          tokenDao,
			userDao:           userDao,
			uuid:              uuid,
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
		roleDao *dao.RoleDao,
		tenantDao *dao.TenantDao,
		userTenantRoleDao *dao.UserTenantRoleDao,
	) (*UserTenantRoleService, error) {
		obj := &UserTenantRoleService{
			roleDao:           roleDao,
			tenantDao:         tenantDao,
			userTenantRoleDao: userTenantRoleDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		hash *bcrypt.Hash,
		hashID *hashids.HashID,
		permissionSvc *PermissionService,
		userDao *dao.UserDao,
		userTenantRoleSvc *UserTenantRoleService,
	) (*UserService, error) {
		obj := &UserService{
			hash:              hash,
			hashID:            hashID,
			permissionSvc:     permissionSvc,
			userDao:           userDao,
			userTenantRoleSvc: userTenantRoleSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
