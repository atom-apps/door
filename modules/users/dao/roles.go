package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/users/dto"

	"gorm.io/gen/field"
)

// @provider
type RoleDao struct {
	query *query.Query
}

func (dao *RoleDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *RoleDao) Context(ctx context.Context) query.IRoleDo {
	return dao.query.Role.WithContext(ctx)
}

func (dao *RoleDao) decorateSortQueryFilter(query query.IRoleDo, sortFilter *ds.SortQueryFilter) query.IRoleDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Role.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Role.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *RoleDao) decorateQueryFilter(query query.IRoleDo, queryFilter *dto.RoleListQueryFilter) query.IRoleDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.Role.Name.Like(common.WrapLike(*queryFilter.Name)))
	}
	if queryFilter.Slug != nil {
		query = query.Where(dao.query.Role.Slug.Like(common.WrapLike(*queryFilter.Slug)))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.Role.Description.Like(common.WrapLike(*queryFilter.Description)))
	}
	if queryFilter.ParentID != nil {
		query = query.Where(dao.query.Role.ParentID.Eq(*queryFilter.ParentID))
	}

	return query
}

func (dao *RoleDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Role.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *RoleDao) Update(ctx context.Context, model *models.Role) error {
	_, err := dao.Context(ctx).Where(dao.query.Role.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *RoleDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Role.ID.Eq(id)).Delete()
	return err
}

func (dao *RoleDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Role.ID.Eq(id)).Delete()
	return err
}

func (dao *RoleDao) Create(ctx context.Context, model *models.Role) error {
	return dao.Context(ctx).Create(model)
}

func (dao *RoleDao) GetByID(ctx context.Context, id uint64) (*models.Role, error) {
	return dao.Context(ctx).Where(dao.query.Role.ID.Eq(id)).First()
}

func (dao *RoleDao) GetBySlug(ctx context.Context, slug string) (*models.Role, error) {
	return dao.Context(ctx).Where(dao.query.Role.Slug.Eq(slug)).First()
}

func (dao *RoleDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Role, error) {
	return dao.Context(ctx).Where(dao.query.Role.ID.In(ids...)).Find()
}

func (dao *RoleDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Role, int64, error) {
	query := dao.query.Role
	roleQuery := query.WithContext(ctx)
	roleQuery = dao.decorateQueryFilter(roleQuery, queryFilter)
	roleQuery = dao.decorateSortQueryFilter(roleQuery, sortFilter)
	return roleQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *RoleDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Role, error) {
	query := dao.query.Role
	roleQuery := query.WithContext(ctx)
	roleQuery = dao.decorateQueryFilter(roleQuery, queryFilter)
	roleQuery = dao.decorateSortQueryFilter(roleQuery, sortFilter)
	return roleQuery.Find()
}

func (dao *RoleDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.Role, error) {
	query := dao.query.Role
	roleQuery := query.WithContext(ctx)
	roleQuery = dao.decorateQueryFilter(roleQuery, queryFilter)
	roleQuery = dao.decorateSortQueryFilter(roleQuery, sortFilter)
	return roleQuery.First()
}

// SlugExists
func (dao *RoleDao) SlugExists(ctx context.Context, model *models.Role) bool {
	count, err := dao.Context(ctx).Where(
		dao.query.Role.Slug.Eq(model.Slug),
		dao.query.Role.ID.Neq(model.ID),
	).Count()
	if err != nil {
		return false
	}
	return count > 0
}

func (dao *RoleDao) FindByIDs(ctx context.Context, ids []uint64) ([]*models.Role, error) {
	table, query := dao.query.Role, dao.Context(ctx)
	return query.Where(table.ID.In(ids...)).Find()
}
