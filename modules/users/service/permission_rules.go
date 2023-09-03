package service

import (
	"context"
	"strconv"

	"github.com/atom-apps/door/common/errorx"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/samber/lo"
)

// @provider
type PermissionRuleService struct {
	dao       *dao.PermissionRuleDao
	tenantDao *dao.TenantDao
	roleDao   *dao.RoleDao
}

func (svc *PermissionRuleService) GetByID(ctx context.Context, id int64) (*models.PermissionRule, error) {
	return svc.dao.GetByID(ctx, id)
}

// role, tenant, path, action
func (svc *PermissionRuleService) genPolicyModel(ctx context.Context, args ...string) *models.PermissionRule {
	return &models.PermissionRule{Ptype: "p", V0: args[0], V1: args[1], V2: args[2], V3: args[3], V4: "", V5: ""}
}

// user, tenant, role
func (svc *PermissionRuleService) genGroupModel(ctx context.Context, args ...string) *models.PermissionRule {
	return &models.PermissionRule{Ptype: "g", V0: args[0], V1: args[1], V2: args[2], V3: "", V4: "", V5: ""}
}

func (svc *PermissionRuleService) GetPolicy(ctx context.Context, tenantID, roleID int64, path, action string) (*models.PermissionRule, error) {
	args := []string{strconv.Itoa(int(roleID)), strconv.Itoa(int(tenantID)), path, action}
	return svc.dao.GetByModel(ctx, svc.genPolicyModel(ctx, args...))
}

// create policy
func (svc *PermissionRuleService) CreatePolicy(ctx context.Context, tenantID, roleID int64, path, action string) error {
	args := []string{strconv.Itoa(int(roleID)), strconv.Itoa(int(tenantID)), path, action}
	if m, err := svc.GetPolicy(ctx, tenantID, roleID, path, action); err == nil && m != nil {
		return errorx.ErrRecordExists
	}
	return svc.dao.Create(ctx, svc.genPolicyModel(ctx, args...))
}

func (svc *PermissionRuleService) DeletePolicy(ctx context.Context, tenantID, roleID int64, path, action string) error {
	args := []string{strconv.Itoa(int(roleID)), strconv.Itoa(int(tenantID)), path, action}
	return svc.dao.DeleteByModel(ctx, svc.genPolicyModel(ctx, args...))
}

// create role
func (svc *PermissionRuleService) CreateGroup(ctx context.Context, userID, tenantID, roleID int64) error {
	args := []string{strconv.Itoa(int(userID)), strconv.Itoa(int(tenantID)), strconv.Itoa(int(roleID))}
	if m, err := svc.GetRoleOfTenantUser(ctx, tenantID, userID); err == nil && m != nil {
		return errorx.ErrRecordExists
	}
	return svc.dao.Create(ctx, svc.genGroupModel(ctx, args...))
}

func (svc *PermissionRuleService) DeleteGroup(ctx context.Context, userID, tenantID, roleID int64) error {
	args := []string{strconv.Itoa(int(userID)), strconv.Itoa(int(tenantID)), strconv.Itoa(int(roleID))}
	return svc.dao.DeleteByModel(ctx, svc.genGroupModel(ctx, args...))
}

func (svc *PermissionRuleService) DeleteUser(ctx context.Context, userID int64) error {
	return svc.dao.DeleteGroupByUser(ctx, userID)
}

func (svc *PermissionRuleService) DeleteRoleUsers(ctx context.Context, tenantID, roleID int64, users []int64) error {
	errs := lo.FilterMap(users, func(userID int64, _ int) (error, bool) {
		if err := svc.DeleteGroup(ctx, userID, tenantID, roleID); err != nil {
			return err, true
		}
		return nil, false
	})

	return lo.Validate(len(errs) == 0, "delete group failed")
}

func (svc *PermissionRuleService) AddRoleUsers(ctx context.Context, tenantID, roleID int64, users []int64) error {
	errs := lo.FilterMap(users, func(userID int64, _ int) (error, bool) {
		if err := svc.CreateGroup(ctx, userID, tenantID, roleID); err != nil {
			return err, true
		}
		return nil, false
	})
	return lo.Validate(len(errs) == 0, "添加失败")
}

func (svc *PermissionRuleService) DeleteByRoleID(ctx context.Context, roleID int64) error {
	if err := svc.dao.DeletePolicyByRoleID(ctx, roleID); err != nil {
		return err
	}

	return svc.dao.DeleteGroupByRoleID(ctx, roleID)
}

func (svc *PermissionRuleService) DeleteByTenantID(ctx context.Context, tenantID int64) error {
	if err := svc.dao.DeletePolicyByTenantID(ctx, tenantID); err != nil {
		return err
	}

	return svc.dao.DeleteGroupByTenantID(ctx, tenantID)
}

func (svc *PermissionRuleService) GetTenantsByUserID(ctx context.Context, userID int64) ([]*models.Tenant, error) {
	tenantIDs, err := svc.GetTenantIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// super admin
	if len(tenantIDs) >= 1 && tenantIDs[0] == 0 {
		return []*models.Tenant{{ID: 0, Name: "Super Admin"}}, nil
	}

	return svc.tenantDao.GetByIDs(ctx, tenantIDs)
}

func (svc *PermissionRuleService) GetTenantIDsByUserID(ctx context.Context, userID int64) ([]int64, error) {
	return svc.dao.GetTenantsByUserID(ctx, userID)
}

func (svc *PermissionRuleService) GetTenantsByRoleID(ctx context.Context, roleID int64) ([]*models.Tenant, error) {
	roleIDs, err := svc.dao.GetTenantsByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	return svc.tenantDao.GetByIDs(ctx, roleIDs)
}

// GetRolesByTenantID
func (svc *PermissionRuleService) GetRolesByTenantID(ctx context.Context, tenantID int64) ([]*models.Role, error) {
	roleIDs, err := svc.dao.GetRolesByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return svc.roleDao.GetByIDs(ctx, roleIDs)
}

// GetRoleOfTenantUser
func (svc *PermissionRuleService) GetRoleOfTenantUser(ctx context.Context, tenantID, userID int64) (*models.Role, error) {
	roleID, err := svc.dao.GetRoleOfTenantUser(ctx, tenantID, userID)
	if err != nil {
		return nil, err
	}

	return svc.roleDao.GetByID(ctx, roleID)
}

// GetUserIDsOfTenant
func (svc *PermissionRuleService) GetUserIDsOfTenant(ctx context.Context, tenantID int64) ([]int64, error) {
	return svc.dao.GetUserIDsOfTenant(ctx, tenantID)
}

func (svc *PermissionRuleService) GetUserAmountOfTenant(ctx context.Context, tenantID int64) (int64, error) {
	return svc.dao.GetUserAmountOfTenant(ctx, tenantID)
}

// GetUserIDsOfRole
func (svc *PermissionRuleService) GetUserIDsOfRole(ctx context.Context, roleID int64) ([]int64, error) {
	return svc.dao.GetUserIDsOfRole(ctx, roleID)
}

func (svc *PermissionRuleService) GetUserAmountOfRole(ctx context.Context, roleID int64) (int64, error) {
	return svc.dao.GetUserAmountOfRole(ctx, roleID)
}
