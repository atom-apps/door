package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/samber/lo"

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

// DeleteByTenantRole
func (dao *PermissionDao) DeleteByTenantRole(ctx context.Context, tenantID uint64, roleID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Permission.TenantID.Eq(tenantID), dao.query.Permission.RoleID.Eq(roleID)).Delete()
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

func (dao *PermissionDao) CreateBatch(ctx context.Context, model []*models.Permission, size int) error {
	return dao.Context(ctx).CreateInBatches(model, size)
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

// GetRouteIDsByTenantIDAndRoleID
func (dao *PermissionDao) GetRouteIDsByTenantIDAndRoleID(ctx context.Context, tenantID, roleID uint64) ([]uint64, error) {
	items, err := dao.Context(ctx).
		Distinct(dao.query.Permission.RouteID).
		Where(
			dao.query.Permission.TenantID.Eq(tenantID),
			dao.query.Permission.RoleID.Eq(roleID),
		).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(items, func(item *models.Permission, _ int) uint64 {
		return item.RouteID
	}), nil
}

func (dao *PermissionDao) FindALl(ctx context.Context) ([]*models.Permission, error) {
	return dao.Context(ctx).Find()
}

func (dao *PermissionDao) FindByTenantRole(ctx context.Context, tenantID, roleID uint64) ([]*models.Permission, error) {
	table, query := dao.query.Permission, dao.Context(ctx)
	return query.Where(
		table.RoleID.Eq(roleID),
		table.TenantID.Eq(tenantID),
	).Find()
}
