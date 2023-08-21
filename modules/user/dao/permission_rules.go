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
