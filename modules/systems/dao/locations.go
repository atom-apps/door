package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/systems/dto"

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
		query = query.Where(dao.query.Location.Code.Like(*queryFilter.Code))
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

// Provinces
func (dao *LocationDao) Provinces(ctx context.Context) ([]*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.City.Eq(""), table.Area.Eq(""), table.Town.Eq("")).Find()
}

// Province
func (dao *LocationDao) Province(ctx context.Context, province string) (*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Province.Eq(province), table.City.Eq(""), table.Area.Eq(""), table.Town.Eq("")).First()
}

// Cities
func (dao *LocationDao) Cities(ctx context.Context, province string) ([]*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Province.Eq(province), table.City.Neq(""), table.Area.Eq(""), table.Town.Eq("")).Find()
}

// City
func (dao *LocationDao) City(ctx context.Context, province, city string) (*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Province.Eq(province), table.City.Eq(city), table.Area.Eq(""), table.Town.Eq("")).First()
}

// Areas
func (dao *LocationDao) Areas(ctx context.Context, province, city string) ([]*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Province.Eq(province), table.City.Eq(city), table.Area.Neq(""), table.Town.Eq("")).Find()
}

// Area
func (dao *LocationDao) Area(ctx context.Context, province, city, area string) (*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Province.Eq(province), table.City.Eq(city), table.Area.Eq(area), table.Town.Eq("")).First()
}

// Towns
func (dao *LocationDao) Towns(ctx context.Context, province, city, area string) ([]*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Province.Eq(province), table.City.Eq(city), table.Area.Eq(area), table.Town.Neq("")).Find()
}

// Town
func (dao *LocationDao) Town(ctx context.Context, province, city, area, town string) (*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Province.Eq(province), table.City.Eq(city), table.Area.Eq(area), table.Town.Eq(town)).First()
}

// GetByCode
func (dao *LocationDao) GetByCode(ctx context.Context, code string) (*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Code.Eq(code)).First()
}

// GetByCode
func (dao *LocationDao) GetByCodeTown(ctx context.Context, code, town string) (*models.Location, error) {
	table, query := dao.query.Location, dao.Context(ctx)
	return query.Where(table.Code.Eq(code), table.Town.Eq(town)).First()
}
