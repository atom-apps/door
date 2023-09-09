package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/samber/lo"

	"gorm.io/gen/field"
)

// @provider
type MenuDao struct {
	query *query.Query
}

func (dao *MenuDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *MenuDao) Context(ctx context.Context) query.IMenuDo {
	return dao.query.Menu.WithContext(ctx)
}

func (dao *MenuDao) decorateSortQueryFilter(query query.IMenuDo, sortFilter *common.SortQueryFilter) query.IMenuDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Menu.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Menu.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *MenuDao) decorateQueryFilter(query query.IMenuDo, queryFilter *dto.MenuListQueryFilter) query.IMenuDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.Menu.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Slug != nil {
		query = query.Where(dao.query.Menu.Slug.Eq(*queryFilter.Slug))
	}
	if queryFilter.GroupID != nil {
		query = query.Where(dao.query.Menu.GroupID.Eq(*queryFilter.GroupID))
	}
	if queryFilter.ParentID != nil {
		query = query.Where(dao.query.Menu.ParentID.Eq(*queryFilter.ParentID))
	}

	return query
}

func (dao *MenuDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Menu.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *MenuDao) Update(ctx context.Context, model *models.Menu) error {
	_, err := dao.Context(ctx).Where(dao.query.Menu.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *MenuDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Menu.ID.Eq(id)).Delete()
	return err
}

func (dao *MenuDao) DeleteAll(ctx context.Context, id, groupID uint64) error {
	_, err := dao.Context(ctx).Where(
		dao.query.Menu.ID.Eq(id),
		dao.query.Menu.GroupID.Eq(groupID),
	).Delete()
	return err
}

func (dao *MenuDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Menu.ID.Eq(id)).Delete()
	return err
}

func (dao *MenuDao) Create(ctx context.Context, model *models.Menu) error {
	return dao.Context(ctx).Create(model)
}

func (dao *MenuDao) GetByID(ctx context.Context, id uint64) (*models.Menu, error) {
	return dao.Context(ctx).Where(dao.query.Menu.ID.Eq(id)).First()
}

func (dao *MenuDao) GetGroupByID(ctx context.Context, groupID uint64) (*models.Menu, error) {
	return dao.Context(ctx).Where(
		dao.query.Menu.ID.Eq(groupID),
		dao.query.Menu.GroupID.Eq(0),
	).First()
}

// GetGroupItemsByID
func (dao *MenuDao) GetGroupItemsByID(ctx context.Context, groupID uint64) ([]*models.Menu, error) {
	ids := []uint64{groupID}

	var items []*models.Menu
	for {
		menus, err := dao.Context(ctx).Where(dao.query.Menu.GroupID.Eq(groupID), dao.query.Menu.ParentID.In(ids...)).Find()
		if err != nil {
			return nil, err
		}

		if len(menus) == 0 {
			break
		}

		ids = lo.Map(menus, func(item *models.Menu, _ int) uint64 {
			return item.ID
		})

		items = append(items, menus...)
	}
	return items, nil
}

func (dao *MenuDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Menu, error) {
	return dao.Context(ctx).Where(dao.query.Menu.ID.In(ids...)).Find()
}

func (dao *MenuDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.MenuListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Menu, int64, error) {
	query := dao.query.Menu
	menuQuery := query.WithContext(ctx)
	menuQuery = dao.decorateQueryFilter(menuQuery, queryFilter)
	menuQuery = dao.decorateSortQueryFilter(menuQuery, sortFilter)
	return menuQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *MenuDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.MenuListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Menu, error) {
	query := dao.query.Menu
	menuQuery := query.WithContext(ctx)
	menuQuery = dao.decorateQueryFilter(menuQuery, queryFilter)
	menuQuery = dao.decorateSortQueryFilter(menuQuery, sortFilter)
	return menuQuery.Find()
}

func (dao *MenuDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.MenuListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Menu, error) {
	query := dao.query.Menu
	menuQuery := query.WithContext(ctx)
	menuQuery = dao.decorateQueryFilter(menuQuery, queryFilter)
	menuQuery = dao.decorateSortQueryFilter(menuQuery, sortFilter)
	return menuQuery.First()
}
