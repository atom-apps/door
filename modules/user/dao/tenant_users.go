package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/user/dto"

	"gorm.io/gen/field"
)

// @provider
type TenantUserDao struct {
	query *query.Query
}

func (dao *TenantUserDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *TenantUserDao) Context(ctx context.Context) query.ITenantUserDo {
	return dao.query.TenantUser.WithContext(ctx)
}

func (dao *TenantUserDao) decorateSortQueryFilter(query query.ITenantUserDo, sortFilter *common.SortQueryFilter) query.ITenantUserDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.TenantUser.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.TenantUser.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *TenantUserDao) decorateQueryFilter(query query.ITenantUserDo, queryFilter *dto.TenantUserListQueryFilter) query.ITenantUserDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.TenantUser.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.TenantUser.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.IsAdmin != nil {
		query = query.Where(dao.query.TenantUser.IsAdmin.Is(*queryFilter.IsAdmin))
	}

	return query
}

func (dao *TenantUserDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.TenantUser.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *TenantUserDao) Update(ctx context.Context, model *models.TenantUser) error {
	_, err := dao.Context(ctx).Where(dao.query.TenantUser.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *TenantUserDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.TenantUser.ID.Eq(id)).Delete()
	return err
}

func (dao *TenantUserDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.TenantUser.ID.Eq(id)).Delete()
	return err
}

func (dao *TenantUserDao) Create(ctx context.Context, model *models.TenantUser) error {
	return dao.Context(ctx).Create(model)
}

func (dao *TenantUserDao) GetByID(ctx context.Context, id int64) (*models.TenantUser, error) {
	return dao.Context(ctx).Where(dao.query.TenantUser.ID.Eq(id)).First()
}

func (dao *TenantUserDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.TenantUser, error) {
	return dao.Context(ctx).Where(dao.query.TenantUser.ID.In(ids...)).Find()
}

func (dao *TenantUserDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantUserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.TenantUser, int64, error) {
	query := dao.query.TenantUser
	tenantUserQuery := query.WithContext(ctx)
	tenantUserQuery = dao.decorateQueryFilter(tenantUserQuery, queryFilter)
	tenantUserQuery = dao.decorateSortQueryFilter(tenantUserQuery, sortFilter)
	return tenantUserQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *TenantUserDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantUserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.TenantUser, error) {
	query := dao.query.TenantUser
	tenantUserQuery := query.WithContext(ctx)
	tenantUserQuery = dao.decorateQueryFilter(tenantUserQuery, queryFilter)
	tenantUserQuery = dao.decorateSortQueryFilter(tenantUserQuery, sortFilter)
	return tenantUserQuery.Find()
}

func (dao *TenantUserDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantUserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.TenantUser, error) {
	query := dao.query.TenantUser
	tenantUserQuery := query.WithContext(ctx)
	tenantUserQuery = dao.decorateQueryFilter(tenantUserQuery, queryFilter)
	tenantUserQuery = dao.decorateSortQueryFilter(tenantUserQuery, sortFilter)
	return tenantUserQuery.First()
}
