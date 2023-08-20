package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/samber/lo"

	"gorm.io/gen/field"
)

// @provider
type RoleUserDao struct {
	query *query.Query
}

func (dao *RoleUserDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *RoleUserDao) Context(ctx context.Context) query.IRoleUserDo {
	return dao.query.RoleUser.WithContext(ctx)
}

func (dao *RoleUserDao) decorateSortQueryFilter(query query.IRoleUserDo, sortFilter *common.SortQueryFilter) query.IRoleUserDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.RoleUser.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.RoleUser.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *RoleUserDao) decorateQueryFilter(query query.IRoleUserDo, queryFilter *dto.RoleUserListQueryFilter) query.IRoleUserDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.RoleID != nil {
		query = query.Where(dao.query.RoleUser.RoleID.Eq(*queryFilter.RoleID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.RoleUser.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.RoleUser.TenantID.Eq(*queryFilter.TenantID))
	}

	return query
}

func (dao *RoleUserDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.RoleUser.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *RoleUserDao) Update(ctx context.Context, model *models.RoleUser) error {
	_, err := dao.Context(ctx).Where(dao.query.RoleUser.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *RoleUserDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.RoleUser.ID.Eq(id)).Delete()
	return err
}

func (dao *RoleUserDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.RoleUser.ID.Eq(id)).Delete()
	return err
}

func (dao *RoleUserDao) Create(ctx context.Context, model *models.RoleUser) error {
	return dao.Context(ctx).Create(model)
}

func (dao *RoleUserDao) GetByID(ctx context.Context, id int64) (*models.RoleUser, error) {
	return dao.Context(ctx).Where(dao.query.RoleUser.ID.Eq(id)).First()
}

func (dao *RoleUserDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.RoleUser, error) {
	return dao.Context(ctx).Where(dao.query.RoleUser.ID.In(ids...)).Find()
}

func (dao *RoleUserDao) GetByUserID(ctx context.Context, tenantID, userID int64) (*models.RoleUser, error) {
	return dao.Context(ctx).Where(
		dao.query.RoleUser.TenantID.Eq(tenantID),
		dao.query.RoleUser.UserID.Eq(userID),
	).First()
}

func (dao *RoleUserDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleUserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.RoleUser, int64, error) {
	query := dao.query.RoleUser
	roleUserQuery := query.WithContext(ctx)
	roleUserQuery = dao.decorateQueryFilter(roleUserQuery, queryFilter)
	roleUserQuery = dao.decorateSortQueryFilter(roleUserQuery, sortFilter)
	return roleUserQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *RoleUserDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleUserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.RoleUser, error) {
	query := dao.query.RoleUser
	roleUserQuery := query.WithContext(ctx)
	roleUserQuery = dao.decorateQueryFilter(roleUserQuery, queryFilter)
	roleUserQuery = dao.decorateSortQueryFilter(roleUserQuery, sortFilter)
	return roleUserQuery.Find()
}

func (dao *RoleUserDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleUserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.RoleUser, error) {
	query := dao.query.RoleUser
	roleUserQuery := query.WithContext(ctx)
	roleUserQuery = dao.decorateQueryFilter(roleUserQuery, queryFilter)
	roleUserQuery = dao.decorateSortQueryFilter(roleUserQuery, sortFilter)
	return roleUserQuery.First()
}

// GetByRoleID
func (dao *RoleUserDao) GetByRoleID(ctx context.Context, roleID int64) ([]*models.RoleUser, error) {
	return dao.Context(ctx).Where(dao.query.RoleUser.RoleID.Eq(roleID)).Find()
}

// AttachUsers
func (dao *RoleUserDao) AttachUsers(ctx context.Context, roleID int64, users []int64) error {
	models := lo.Map(users, func(id int64, _ int) *models.RoleUser {
		return &models.RoleUser{
			RoleID: roleID,
			UserID: id,
		}
	})

	return dao.Context(ctx).CreateInBatches(models, 10)
}

// DetachUsers
func (dao *RoleUserDao) DetachUsers(ctx context.Context, roleID int64, users []int64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.RoleUser.RoleID.Eq(roleID),
		dao.query.RoleUser.UserID.In(users...),
	).Delete()
	return err
}
