package dao

import (
	"context"
	"strconv"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/samber/lo"
)

// @provider
type PermissionRuleDao struct {
	query *query.Query
}

func (dao *PermissionRuleDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *PermissionRuleDao) Context(ctx context.Context) query.IPermissionRuleDo {
	return dao.query.PermissionRule.WithContext(ctx)
}

func (dao *PermissionRuleDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.PermissionRule.ID.Eq(id)).Delete()
	return err
}

func (dao *PermissionRuleDao) GetByModel(ctx context.Context, m *models.PermissionRule) (*models.PermissionRule, error) {
	return dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq(m.Ptype),
		dao.query.PermissionRule.V0.Eq(m.V0),
		dao.query.PermissionRule.V1.Eq(m.V1),
		dao.query.PermissionRule.V2.Eq(m.V2),
		dao.query.PermissionRule.V3.Eq(m.V3),
		dao.query.PermissionRule.V4.Eq(m.V4),
		dao.query.PermissionRule.V5.Eq(m.V5),
	).First()
}

func (dao *PermissionRuleDao) DeleteByModel(ctx context.Context, m *models.PermissionRule) error {
	_, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq(m.Ptype),
		dao.query.PermissionRule.V0.Eq(m.V0),
		dao.query.PermissionRule.V1.Eq(m.V1),
		dao.query.PermissionRule.V2.Eq(m.V2),
		dao.query.PermissionRule.V3.Eq(m.V3),
		dao.query.PermissionRule.V4.Eq(m.V4),
		dao.query.PermissionRule.V5.Eq(m.V5),
	).Delete()
	return err
}

func (dao *PermissionRuleDao) DeleteGroupByUser(ctx context.Context, userID int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("p"),
		dao.query.PermissionRule.V0.Eq(strconv.Itoa(int(userID))),
	).Delete()
	return err
}

func (dao *PermissionRuleDao) Create(ctx context.Context, model *models.PermissionRule) error {
	return dao.Context(ctx).Create(model)
}

func (dao *PermissionRuleDao) GetByID(ctx context.Context, id int64) (*models.PermissionRule, error) {
	return dao.Context(ctx).Where(dao.query.PermissionRule.ID.Eq(id)).First()
}

// DeleteRoleUsers
func (dao *PermissionRuleDao) DeleteRoleUsers(ctx context.Context, tenantID, roleID int64, users []int64) error {
	userStringIDs := lo.Map(users, func(user int64, _ int) string {
		return strconv.Itoa(int(user))
	})

	_, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V0.In(userStringIDs...),
		dao.query.PermissionRule.V1.Eq(strconv.Itoa(int(roleID))),
		dao.query.PermissionRule.V2.Eq(strconv.Itoa(int(tenantID))),
	).Delete()
	return err
}

// DeleteRoleUsers
func (dao *PermissionRuleDao) DeletePolicyByRoleID(ctx context.Context, roleID int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("p"),
		dao.query.PermissionRule.V0.Eq(strconv.Itoa(int(roleID))),
	).Delete()
	return err
}

// DeleteGroupByRoleID
func (dao *PermissionRuleDao) DeleteGroupByRoleID(ctx context.Context, roleID int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V1.Eq(strconv.Itoa(int(roleID))),
	).Delete()
	return err
}

// DeleteRoleUsers
func (dao *PermissionRuleDao) DeletePolicyByTenantID(ctx context.Context, tenantID int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("p"),
		dao.query.PermissionRule.V1.Eq(strconv.Itoa(int(tenantID))),
	).Delete()
	return err
}

// DeleteGroupByTenantID
func (dao *PermissionRuleDao) DeleteGroupByTenantID(ctx context.Context, tenantID int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V2.Eq(strconv.Itoa(int(tenantID))),
	).Delete()
	return err
}

// GetTenantsByUserID
func (dao *PermissionRuleDao) GetTenantsByUserID(ctx context.Context, userID int64) ([]int64, error) {
	ms, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V0.Eq(strconv.Itoa(int(userID))),
	).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(ms, func(m *models.PermissionRule, _ int) int64 {
		tenantID, _ := strconv.Atoi(m.V1)
		return int64(tenantID)
	}), nil
}

// GetTenantsByRoleID
func (dao *PermissionRuleDao) GetTenantsByRoleID(ctx context.Context, roleID int64) ([]int64, error) {
	ms, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V2.Eq(strconv.Itoa(int(roleID))),
	).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(ms, func(m *models.PermissionRule, _ int) int64 {
		tenantID, _ := strconv.Atoi(m.V1)
		return int64(tenantID)
	}), nil
}

// GetRolesByTenantID
func (dao *PermissionRuleDao) GetRolesByTenantID(ctx context.Context, tenantID int64) ([]int64, error) {
	ms, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V1.Eq(strconv.Itoa(int(tenantID))),
	).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(ms, func(m *models.PermissionRule, _ int) int64 {
		tenantID, _ := strconv.Atoi(m.V2)
		return int64(tenantID)
	}), nil
}

// GetRoleOfTenantUser
func (dao *PermissionRuleDao) GetRoleOfTenantUser(ctx context.Context, tenantID, userID int64) (int64, error) {
	m, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V0.Eq(strconv.Itoa(int(userID))),
		dao.query.PermissionRule.V1.Eq(strconv.Itoa(int(tenantID))),
	).First()
	if err != nil {
		return 0, err
	}
	tid, err := strconv.Atoi(m.V2)
	return int64(tid), err
}

// GetUserIDsOfTenant
func (dao *PermissionRuleDao) GetUserIDsOfTenant(ctx context.Context, tenantID int64) ([]int64, error) {
	ms, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V1.Eq(strconv.Itoa(int(tenantID))),
	).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(ms, func(m *models.PermissionRule, _ int) int64 {
		id, _ := strconv.Atoi(m.V0)
		return int64(id)
	}), nil
}

func (dao *PermissionRuleDao) GetUserAmountOfTenant(ctx context.Context, tenantID int64) (int64, error) {
	return dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V1.Eq(strconv.Itoa(int(tenantID))),
	).Count()
}

// GetUserIDsOfRole
func (dao *PermissionRuleDao) GetUserIDsOfRole(ctx context.Context, roleID int64) ([]int64, error) {
	ms, err := dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V2.Eq(strconv.Itoa(int(roleID))),
	).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(ms, func(m *models.PermissionRule, _ int) int64 {
		id, _ := strconv.Atoi(m.V0)
		return int64(id)
	}), nil
}

func (dao *PermissionRuleDao) GetUserAmountOfRole(ctx context.Context, roleID int64) (int64, error) {
	return dao.Context(ctx).Where(
		dao.query.PermissionRule.Ptype.Eq("g"),
		dao.query.PermissionRule.V2.Eq(strconv.Itoa(int(roleID))),
	).Count()
}
