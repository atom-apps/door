package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/users/dto"

	"gorm.io/gen/field"
)

// @provider
type PermissionDao struct {
	query *query.Query
}

func (dao *PermissionDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *PermissionDao) Context(ctx context.Context) query.IPermissionDo {
	return dao.query.Permission.WithContext(ctx)
}

func (dao *PermissionDao) decorateSortQueryFilter(query query.IPermissionDo, sortFilter *common.SortQueryFilter) query.IPermissionDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Permission.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Permission.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *PermissionDao) decorateQueryFilter(query query.IPermissionDo, queryFilter *dto.PermissionListQueryFilter) query.IPermissionDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.Permission.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.RoleID != nil {
		query = query.Where(dao.query.Permission.RoleID.Eq(*queryFilter.RoleID))
	}
	if queryFilter.Path != nil {
		query = query.Where(dao.query.Permission.Path.Eq(*queryFilter.Path))
	}
	if queryFilter.Action != nil {
		query = query.Where(dao.query.Permission.Action.Eq(*queryFilter.Action))
	}

	return query
}

func (dao *PermissionDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Permission.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *PermissionDao) Update(ctx context.Context, model *models.Permission) error {
	_, err := dao.Context(ctx).Where(dao.query.Permission.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *PermissionDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Permission.ID.Eq(id)).Delete()
	return err
}

func (dao *PermissionDao) DeleteByTenantID(ctx context.Context, tenantID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Permission.TenantID.Eq(tenantID)).Delete()
	return err
}

func (dao *PermissionDao) DeleteByRoleID(ctx context.Context, roleID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Permission.RoleID.Eq(roleID)).Delete()
	return err
}

func (dao *PermissionDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Permission.ID.Eq(id)).Delete()
	return err
}

func (dao *PermissionDao) Create(ctx context.Context, model *models.Permission) error {
	return dao.Context(ctx).Create(model)
}

func (dao *PermissionDao) GetByID(ctx context.Context, id uint64) (*models.Permission, error) {
	return dao.Context(ctx).Where(dao.query.Permission.ID.Eq(id)).First()
}

func (dao *PermissionDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Permission, error) {
	return dao.Context(ctx).Where(dao.query.Permission.ID.In(ids...)).Find()
}

func (dao *PermissionDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.PermissionListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Permission, int64, error) {
	query := dao.query.Permission
	permissionQuery := query.WithContext(ctx)
	permissionQuery = dao.decorateQueryFilter(permissionQuery, queryFilter)
	permissionQuery = dao.decorateSortQueryFilter(permissionQuery, sortFilter)
	return permissionQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *PermissionDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.PermissionListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Permission, error) {
	query := dao.query.Permission
	permissionQuery := query.WithContext(ctx)
	permissionQuery = dao.decorateQueryFilter(permissionQuery, queryFilter)
	permissionQuery = dao.decorateSortQueryFilter(permissionQuery, sortFilter)
	return permissionQuery.Find()
}

func (dao *PermissionDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.PermissionListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Permission, error) {
	query := dao.query.Permission
	permissionQuery := query.WithContext(ctx)
	permissionQuery = dao.decorateQueryFilter(permissionQuery, queryFilter)
	permissionQuery = dao.decorateSortQueryFilter(permissionQuery, sortFilter)
	return permissionQuery.First()
}
