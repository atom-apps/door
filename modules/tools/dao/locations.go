package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/tools/dto"

	"gorm.io/gen/field"
)

// @provider
type LocationDao struct {
	query *query.Query
}

func (dao *LocationDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *LocationDao) Context(ctx context.Context) query.ILocationDo {
	return dao.query.Location.WithContext(ctx)
}

func (dao *LocationDao) decorateSortQueryFilter(query query.ILocationDo, sortFilter *common.SortQueryFilter) query.ILocationDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Location.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Location.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *LocationDao) decorateQueryFilter(query query.ILocationDo, queryFilter *dto.LocationListQueryFilter) query.ILocationDo {
	if queryFilter == nil {
		return query
	}

	if queryFilter.Code != nil {
		query = query.Where(dao.query.Location.Code.Eq(*queryFilter.Code))
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.Location.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Province != nil {
		query = query.Where(dao.query.Location.Province.Eq(*queryFilter.Province))
	}
	if queryFilter.City != nil {
		query = query.Where(dao.query.Location.City.Eq(*queryFilter.City))
	}
	if queryFilter.Area != nil {
		query = query.Where(dao.query.Location.Area.Eq(*queryFilter.Area))
	}
	if queryFilter.Town != nil {
		query = query.Where(dao.query.Location.Town.Eq(*queryFilter.Town))
	}

	return query
}

func (dao *LocationDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Location.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *LocationDao) Update(ctx context.Context, model *models.Location) error {
	_, err := dao.Context(ctx).Where(dao.query.Location.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *LocationDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Location.ID.Eq(id)).Delete()
	return err
}

func (dao *LocationDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Location.ID.Eq(id)).Delete()
	return err
}

func (dao *LocationDao) Create(ctx context.Context, model *models.Location) error {
	return dao.Context(ctx).Create(model)
}

func (dao *LocationDao) GetByID(ctx context.Context, id uint64) (*models.Location, error) {
	return dao.Context(ctx).Where(dao.query.Location.ID.Eq(id)).First()
}

func (dao *LocationDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Location, error) {
	return dao.Context(ctx).Where(dao.query.Location.ID.In(ids...)).Find()
}

func (dao *LocationDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.LocationListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Location, int64, error) {
	query := dao.query.Location
	locationQuery := query.WithContext(ctx)
	locationQuery = dao.decorateQueryFilter(locationQuery, queryFilter)
	locationQuery = dao.decorateSortQueryFilter(locationQuery, sortFilter)
	return locationQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *LocationDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.LocationListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Location, error) {
	query := dao.query.Location
	locationQuery := query.WithContext(ctx)
	locationQuery = dao.decorateQueryFilter(locationQuery, queryFilter)
	locationQuery = dao.decorateSortQueryFilter(locationQuery, sortFilter)
	return locationQuery.Find()
}

func (dao *LocationDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.LocationListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Location, error) {
	query := dao.query.Location
	locationQuery := query.WithContext(ctx)
	locationQuery = dao.decorateQueryFilter(locationQuery, queryFilter)
	locationQuery = dao.decorateSortQueryFilter(locationQuery, sortFilter)
	return locationQuery.First()
}
