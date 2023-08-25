package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/systems/dto"

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

func (dao *RouteDao) decorateSortQueryFilter(query query.IRouteDo, sortFilter *common.SortQueryFilter) query.IRouteDo {
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
	if queryFilter.Type != nil {
		query = query.Where(dao.query.Route.Type.Eq(*queryFilter.Type))
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

func (dao *RouteDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Route.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *RouteDao) Update(ctx context.Context, model *models.Route) error {
	_, err := dao.Context(ctx).Where(dao.query.Route.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *RouteDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.Route.ID.Eq(id)).Delete()
	return err
}

func (dao *RouteDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Route.ID.Eq(id)).Delete()
	return err
}

func (dao *RouteDao) Create(ctx context.Context, model *models.Route) error {
	return dao.Context(ctx).Create(model)
}

func (dao *RouteDao) GetByID(ctx context.Context, id int64) (*models.Route, error) {
	return dao.Context(ctx).Where(dao.query.Route.ID.Eq(id)).First()
}

func (dao *RouteDao) FindByParentIDOfMode(ctx context.Context, mode consts.RouteType, parentID int64) ([]*models.Route, error) {
	return dao.Context(ctx).Where(
		dao.query.Route.ParentID.Eq(parentID),
		dao.query.Route.Type.Eq(mode),
	).Find()
}

func (dao *RouteDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.Route, error) {
	return dao.Context(ctx).Where(dao.query.Route.ID.In(ids...)).Find()
}

func (dao *RouteDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RouteListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
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
	sortFilter *common.SortQueryFilter,
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
	sortFilter *common.SortQueryFilter,
) (*models.Route, error) {
	query := dao.query.Route
	routeQuery := query.WithContext(ctx)
	routeQuery = dao.decorateQueryFilter(routeQuery, queryFilter)
	routeQuery = dao.decorateSortQueryFilter(routeQuery, sortFilter)
	return routeQuery.First()
}
