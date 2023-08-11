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
type TenantDao struct {
	query *query.Query
}

func (dao *TenantDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *TenantDao) Context(ctx context.Context) query.ITenantDo {
	return dao.query.Tenant.WithContext(ctx)
}

func (dao *TenantDao) decorateSortQueryFilter(query query.ITenantDo, sortFilter *common.SortQueryFilter) query.ITenantDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Tenant.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Tenant.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *TenantDao) decorateQueryFilter(query query.ITenantDo, queryFilter *dto.TenantListQueryFilter) query.ITenantDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.Tenant.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.Tenant.Description.Eq(*queryFilter.Description))
	}
	if queryFilter.Meta != nil {
		query = query.Where(dao.query.Tenant.Meta.Eq(*queryFilter.Meta))
	}

	return query
}

func (dao *TenantDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Tenant.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *TenantDao) Update(ctx context.Context, model *models.Tenant) error {
	_, err := dao.Context(ctx).Where(dao.query.Tenant.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *TenantDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.Tenant.ID.Eq(id)).Delete()
	return err
}

func (dao *TenantDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Tenant.ID.Eq(id)).Delete()
	return err
}

func (dao *TenantDao) Restore(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Tenant.ID.Eq(id)).UpdateSimple(dao.query.Tenant.DeletedAt.Null())
	return err
}

func (dao *TenantDao) Create(ctx context.Context, model *models.Tenant) error {
	return dao.Context(ctx).Create(model)
}

func (dao *TenantDao) GetByID(ctx context.Context, id int64) (*models.Tenant, error) {
	return dao.Context(ctx).Where(dao.query.Tenant.ID.Eq(id)).First()
}

func (dao *TenantDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.Tenant, error) {
	return dao.Context(ctx).Where(dao.query.Tenant.ID.In(ids...)).Find()
}

func (dao *TenantDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Tenant, int64, error) {
	query := dao.query.Tenant
	tenantQuery := query.WithContext(ctx)
	tenantQuery = dao.decorateQueryFilter(tenantQuery, queryFilter)
	tenantQuery = dao.decorateSortQueryFilter(tenantQuery, sortFilter)
	return tenantQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *TenantDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Tenant, error) {
	query := dao.query.Tenant
	tenantQuery := query.WithContext(ctx)
	tenantQuery = dao.decorateQueryFilter(tenantQuery, queryFilter)
	tenantQuery = dao.decorateSortQueryFilter(tenantQuery, sortFilter)
	return tenantQuery.Find()
}

func (dao *TenantDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Tenant, error) {
	query := dao.query.Tenant
	tenantQuery := query.WithContext(ctx)
	tenantQuery = dao.decorateQueryFilter(tenantQuery, queryFilter)
	tenantQuery = dao.decorateSortQueryFilter(tenantQuery, sortFilter)
	return tenantQuery.First()
}
