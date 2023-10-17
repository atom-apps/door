package dao

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/samber/lo"

	"gorm.io/gen/field"
)

// @provider
type UserTenantRoleDao struct {
	query *query.Query
}

func (dao *UserTenantRoleDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *UserTenantRoleDao) Context(ctx context.Context) query.IUserTenantRoleDo {
	return dao.query.UserTenantRole.WithContext(ctx)
}

func (dao *UserTenantRoleDao) decorateSortQueryFilter(query query.IUserTenantRoleDo, sortFilter *ds.SortQueryFilter) query.IUserTenantRoleDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.UserTenantRole.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.UserTenantRole.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *UserTenantRoleDao) decorateQueryFilter(query query.IUserTenantRoleDo, queryFilter *dto.UserTenantRoleListQueryFilter) query.IUserTenantRoleDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.UserTenantRole.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.UserTenantRole.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.RoleID != nil {
		query = query.Where(dao.query.UserTenantRole.RoleID.Eq(*queryFilter.RoleID))
	}

	return query
}

func (dao *UserTenantRoleDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.UserTenantRole.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *UserTenantRoleDao) Update(ctx context.Context, model *models.UserTenantRole) error {
	_, err := dao.Context(ctx).Where(dao.query.UserTenantRole.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *UserTenantRoleDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.UserTenantRole.ID.Eq(id)).Delete()
	return err
}

func (dao *UserTenantRoleDao) DeleteByRoleID(ctx context.Context, roleID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.UserTenantRole.RoleID.Eq(roleID)).Delete()
	return err
}

func (dao *UserTenantRoleDao) DeleteByTenantID(ctx context.Context, tenantID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.UserTenantRole.TenantID.Eq(tenantID)).Delete()
	return err
}

func (dao *UserTenantRoleDao) DeleteInBatch(ctx context.Context, ms []*models.UserTenantRole) error {
	lo.ForEach(ms, func(item *models.UserTenantRole, _ int) {
		dao.Context(ctx).Where(
			dao.query.UserTenantRole.UserID.Eq(item.UserID),
			dao.query.UserTenantRole.TenantID.Eq(item.TenantID),
			dao.query.UserTenantRole.RoleID.Eq(item.RoleID),
		).Delete()
	})
	return nil
}

func (dao *UserTenantRoleDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.UserTenantRole.ID.Eq(id)).Delete()
	return err
}

func (dao *UserTenantRoleDao) Create(ctx context.Context, model *models.UserTenantRole) error {
	return dao.Context(ctx).Create(model)
}

func (dao *UserTenantRoleDao) CreateInBatch(ctx context.Context, models []*models.UserTenantRole) error {
	return dao.Context(ctx).CreateInBatches(models, 10)
}

func (dao *UserTenantRoleDao) GetByID(ctx context.Context, id uint64) (*models.UserTenantRole, error) {
	return dao.Context(ctx).Where(dao.query.UserTenantRole.ID.Eq(id)).First()
}

func (dao *UserTenantRoleDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.UserTenantRole, error) {
	return dao.Context(ctx).Where(dao.query.UserTenantRole.ID.In(ids...)).Find()
}

func (dao *UserTenantRoleDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserTenantRoleListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.UserTenantRole, int64, error) {
	query := dao.query.UserTenantRole
	userTenantRoleQuery := query.WithContext(ctx)
	userTenantRoleQuery = dao.decorateQueryFilter(userTenantRoleQuery, queryFilter)
	userTenantRoleQuery = dao.decorateSortQueryFilter(userTenantRoleQuery, sortFilter)
	return userTenantRoleQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *UserTenantRoleDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserTenantRoleListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.UserTenantRole, error) {
	query := dao.query.UserTenantRole
	userTenantRoleQuery := query.WithContext(ctx)
	userTenantRoleQuery = dao.decorateQueryFilter(userTenantRoleQuery, queryFilter)
	userTenantRoleQuery = dao.decorateSortQueryFilter(userTenantRoleQuery, sortFilter)
	return userTenantRoleQuery.Find()
}

func (dao *UserTenantRoleDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserTenantRoleListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.UserTenantRole, error) {
	query := dao.query.UserTenantRole
	userTenantRoleQuery := query.WithContext(ctx)
	userTenantRoleQuery = dao.decorateQueryFilter(userTenantRoleQuery, queryFilter)
	userTenantRoleQuery = dao.decorateSortQueryFilter(userTenantRoleQuery, sortFilter)
	return userTenantRoleQuery.First()
}

// GetUserAmountOfRole
func (dao *UserTenantRoleDao) GetUserAmountOfRole(ctx context.Context, roleID uint64) (int64, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)

	return query.Distinct(table.ID).Where(table.RoleID.Eq(roleID)).Count()
}

// GetUserAmountOfTenant
func (dao *UserTenantRoleDao) GetUserAmountOfTenant(ctx context.Context, tenantID uint64) (int64, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)

	return query.Distinct(table.ID).Where(table.TenantID.Eq(tenantID)).Count()
}

// GetRoleOfTenantUser
func (dao *UserTenantRoleDao) GetRoleOfTenantUser(ctx context.Context, tenantID, userID uint64) (*models.UserTenantRole, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)
	return query.Select(table.RoleID).Where(table.TenantID.Eq(tenantID), table.UserID.Eq(userID)).First()
}

// GetTenantsByUserID
func (dao *UserTenantRoleDao) GetTenantsByUserID(ctx context.Context, userID uint64) ([]*models.UserTenantRole, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)
	return query.Distinct(table.UserID).Select(table.TenantID).Where(table.UserID.Eq(userID)).Find()
}

// DeleteByUserID
func (dao *UserTenantRoleDao) DeleteByUserID(ctx context.Context, userID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.UserTenantRole.UserID.Eq(userID)).Delete()
	return err
}

// TenantHasRole
func (dao *UserTenantRoleDao) TenantHasRole(ctx context.Context, tenantID, userID uint64) (bool, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)

	count, err := query.Where(table.TenantID.Eq(tenantID), table.UserID.Eq(userID)).Count()
	if err != nil {
		return true, err
	}

	return count > 0, nil
}

// GetTenantsByRoleID
func (dao *UserTenantRoleDao) GetTenantsByRoleID(ctx context.Context, roleID uint64) ([]*models.UserTenantRole, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)
	return query.Distinct(table.TenantID).Select(table.TenantID).Where(table.RoleID.Eq(roleID)).Find()
}

// GetRolesByTenantID
func (dao *UserTenantRoleDao) GetRolesByTenantID(ctx context.Context, tenantID uint64) ([]*models.UserTenantRole, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)
	return query.Distinct(table.RoleID).Select(table.RoleID).Where(table.TenantID.Eq(tenantID)).Find()
}

// FindAll
func (dao *UserTenantRoleDao) FindAll(ctx context.Context) ([]*models.UserTenantRole, error) {
	return dao.Context(ctx).Find()
}

// FindAll
func (dao *UserTenantRoleDao) FindByUserID(ctx context.Context, userID uint64) ([]*models.UserTenantRole, error) {
	table, query := dao.query.UserTenantRole, dao.Context(ctx)
	return query.Where(table.UserID.Eq(userID)).Find()
}
