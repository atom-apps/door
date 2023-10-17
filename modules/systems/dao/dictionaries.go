package dao

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/systems/dto"

	"gorm.io/gen/field"
)

// @provider
type DictionaryDao struct {
	query *query.Query
}

func (dao *DictionaryDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *DictionaryDao) Context(ctx context.Context) query.IDictionaryDo {
	return dao.query.Dictionary.WithContext(ctx)
}

func (dao *DictionaryDao) decorateSortQueryFilter(query query.IDictionaryDo, sortFilter *ds.SortQueryFilter) query.IDictionaryDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Dictionary.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Dictionary.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *DictionaryDao) decorateQueryFilter(query query.IDictionaryDo, queryFilter *dto.DictionaryListQueryFilter) query.IDictionaryDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.Dictionary.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Slug != nil {
		query = query.Where(dao.query.Dictionary.Slug.Eq(*queryFilter.Slug))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.Dictionary.Description.Eq(*queryFilter.Description))
	}

	return query
}

func (dao *DictionaryDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Dictionary.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *DictionaryDao) Update(ctx context.Context, model *models.Dictionary) error {
	_, err := dao.Context(ctx).Where(dao.query.Dictionary.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *DictionaryDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Dictionary.ID.Eq(id)).Delete()
	return err
}

func (dao *DictionaryDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Dictionary.ID.Eq(id)).Delete()
	return err
}

func (dao *DictionaryDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Dictionary.ID.Eq(id)).UpdateSimple(dao.query.Dictionary.DeletedAt.Null())
	return err
}

func (dao *DictionaryDao) Create(ctx context.Context, model *models.Dictionary) error {
	return dao.Context(ctx).Create(model)
}

func (dao *DictionaryDao) GetByID(ctx context.Context, id uint64) (*models.Dictionary, error) {
	return dao.Context(ctx).Where(dao.query.Dictionary.ID.Eq(id)).First()
}

func (dao *DictionaryDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Dictionary, error) {
	return dao.Context(ctx).Where(dao.query.Dictionary.ID.In(ids...)).Find()
}

func (dao *DictionaryDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Dictionary, int64, error) {
	query := dao.query.Dictionary
	dictionaryQuery := query.WithContext(ctx)
	dictionaryQuery = dao.decorateQueryFilter(dictionaryQuery, queryFilter)
	dictionaryQuery = dao.decorateSortQueryFilter(dictionaryQuery, sortFilter)
	return dictionaryQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *DictionaryDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Dictionary, error) {
	query := dao.query.Dictionary
	dictionaryQuery := query.WithContext(ctx)
	dictionaryQuery = dao.decorateQueryFilter(dictionaryQuery, queryFilter)
	dictionaryQuery = dao.decorateSortQueryFilter(dictionaryQuery, sortFilter)
	return dictionaryQuery.Find()
}

func (dao *DictionaryDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.Dictionary, error) {
	query := dao.query.Dictionary
	dictionaryQuery := query.WithContext(ctx)
	dictionaryQuery = dao.decorateQueryFilter(dictionaryQuery, queryFilter)
	dictionaryQuery = dao.decorateSortQueryFilter(dictionaryQuery, sortFilter)
	return dictionaryQuery.First()
}
