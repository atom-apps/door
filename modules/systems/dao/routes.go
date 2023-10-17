package dao

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/samber/lo"

	"gorm.io/gen/field"
)

// @provider
type RouteDao struct {
	query *query.Query
}

func (dao *RouteDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *RouteDao) Context(ctx context.Context) query.IRouteDo {
	return dao.query.Route.WithContext(ctx)
}

func (dao *RouteDao) decorateSortQueryFilter(query query.IRouteDo, sortFilter *ds.SortQueryFilter) query.IRouteDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Route.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Route.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *RouteDao) decorateQueryFilter(query query.IRouteDo, queryFilter *dto.RouteListQueryFilter) query.IRouteDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.ParentID != nil {
		query = query.Where(dao.query.Route.ParentID.Eq(*queryFilter.ParentID))
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.Route.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Path != nil {
		query = query.Where(dao.query.Route.Path.Eq(*queryFilter.Path))
	}

	return query
}

func (dao *RouteDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Route.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *RouteDao) Update(ctx context.Context, model *models.Route) error {
	_, err := dao.Context(ctx).Where(dao.query.Route.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *RouteDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Route.ID.Eq(id)).Delete()
	return err
}

func (dao *RouteDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Route.ID.Eq(id)).Delete()
	return err
}

func (dao *RouteDao) Create(ctx context.Context, model *models.Route) error {
	return dao.Context(ctx).Create(model)
}

func (dao *RouteDao) GetByID(ctx context.Context, id uint64) (*models.Route, error) {
	return dao.Context(ctx).Where(dao.query.Route.ID.Eq(id)).First()
}

func (dao *RouteDao) FindByParentIDOfMode(ctx context.Context, parentID uint64) ([]*models.Route, error) {
	return dao.Context(ctx).Where(dao.query.Route.ParentID.Eq(parentID)).Order(dao.query.Route.Order, dao.query.Route.ID).Find()
}

func (dao *RouteDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Route, error) {
	return dao.Context(ctx).Where(dao.query.Route.ID.In(ids...)).Find()
}

func (dao *RouteDao) GetByIDsWithParents(ctx context.Context, ids []uint64) ([]*models.Route, error) {
	table, query := dao.query.Route, dao.Context(ctx)

	finalRoutes := []*models.Route{}
	// 如果当前ID是父级ID，那么他的子级权限也应该包含在内
	allIDs := []uint64{}

	// find children
	childrenIds := append([]uint64{}, ids...)
	for {
		children, err := query.Where(table.ParentID.In(childrenIds...)).Find()
		if err != nil {
			return nil, err
		}

		if len(children) == 0 {
			break
		}

		childrenIds = []uint64{}
		lo.ForEach(children, func(item *models.Route, _ int) {
			if lo.Contains(allIDs, item.ID) {
				return
			}
			allIDs = append(allIDs, item.ID)
			childrenIds = append(childrenIds, item.ID)
			finalRoutes = append(finalRoutes, item)
		})
	}

	parentIDs := append([]uint64{}, ids...)
	for len(parentIDs) > 0 {
		routes, err := query.Where(table.ID.In(parentIDs...)).Find()
		if err != nil {
			return nil, err
		}
		finalRoutes = append(finalRoutes, routes...)

		parentIDs = lo.FilterMap(routes, func(item *models.Route, index int) (uint64, bool) {
			if !lo.Contains(allIDs, item.ID) {
				allIDs = append(allIDs, item.ID)
			}

			if item.ParentID != 0 {
				return item.ParentID, true
			}
			return 0, false
		})

		parentIDs = lo.Filter(parentIDs, func(item uint64, index int) bool {
			return !lo.Contains(allIDs, item)
		})
	}

	return finalRoutes, nil
}

func (dao *RouteDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RouteListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Route, int64, error) {
	query := dao.query.Route
	routeQuery := query.WithContext(ctx)
	routeQuery = dao.decorateQueryFilter(routeQuery, queryFilter)
	routeQuery = dao.decorateSortQueryFilter(routeQuery, sortFilter)
	return routeQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *RouteDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RouteListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Route, error) {
	query := dao.query.Route
	routeQuery := query.WithContext(ctx)
	routeQuery = dao.decorateQueryFilter(routeQuery, queryFilter)
	routeQuery = dao.decorateSortQueryFilter(routeQuery, sortFilter)
	return routeQuery.Find()
}

func (dao *RouteDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RouteListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.Route, error) {
	query := dao.query.Route
	routeQuery := query.WithContext(ctx)
	routeQuery = dao.decorateQueryFilter(routeQuery, queryFilter)
	routeQuery = dao.decorateSortQueryFilter(routeQuery, sortFilter)
	return routeQuery.First()
}

func (dao *RouteDao) FindAll(ctx context.Context) ([]*models.Route, error) {
	query := dao.query.Route
	return query.WithContext(ctx).Find()
}
