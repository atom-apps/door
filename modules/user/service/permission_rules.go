package service

import (
	"context"
	"strconv"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/user/dao"
)

// @provider
type PermissionRuleService struct {
	dao *dao.PermissionRuleDao
}

func (svc *PermissionRuleService) GetByID(ctx context.Context, id int64) (*models.PermissionRule, error) {
	return svc.dao.GetByID(ctx, id)
}

func (svc *PermissionRuleService) genPolicyModel(ctx context.Context, args ...string) *models.PermissionRule {
	return &models.PermissionRule{Ptype: "p", V0: args[0], V1: args[1], V2: args[2], V3: args[3], V4: "", V5: ""}
}

func (svc *PermissionRuleService) genRoleModel(ctx context.Context, args ...string) *models.PermissionRule {
	return &models.PermissionRule{Ptype: "g", V0: args[0], V1: args[1], V2: args[2], V3: "", V4: "", V5: ""}
}

// create policy
func (svc *PermissionRuleService) CreatePolicy(ctx context.Context, tenantID, roleID int64, path, action string) error {
	args := []string{strconv.Itoa(int(roleID)), strconv.Itoa(int(tenantID)), path, action}
	return svc.dao.Create(ctx, svc.genPolicyModel(ctx, args...))
}

func (svc *PermissionRuleService) DeletePolicy(ctx context.Context, tenantID, roleID int64, path, action string) error {
	args := []string{strconv.Itoa(int(roleID)), strconv.Itoa(int(tenantID)), path, action}
	return svc.dao.DeleteByModel(ctx, svc.genPolicyModel(ctx, args...))
}

// create role
func (svc *PermissionRuleService) CreateRole(ctx context.Context, userID, tenantID, roleID int64) error {
	args := []string{strconv.Itoa(int(userID)), strconv.Itoa(int(roleID)), strconv.Itoa(int(tenantID))}
	return svc.dao.Create(ctx, svc.genRoleModel(ctx, args...))
}

func (svc *PermissionRuleService) DeleteRole(ctx context.Context, userID, tenantID, roleID int64) error {
	args := []string{strconv.Itoa(int(userID)), strconv.Itoa(int(roleID)), strconv.Itoa(int(tenantID))}
	return svc.dao.DeleteByModel(ctx, svc.genRoleModel(ctx, args...))
}
