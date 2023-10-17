package dao

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/posts/dto"
	"gorm.io/gen/field"
)

// @provider
type BookDao struct {
	query *query.Query
}

func (dao *BookDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *BookDao) Context(ctx context.Context) query.IBookDo {
	return dao.query.Book.WithContext(ctx)
}

func (dao *BookDao) decorateSortQueryFilter(query query.IBookDo, sortFilter *ds.SortQueryFilter) query.IBookDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Book.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Book.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *BookDao) decorateQueryFilter(query query.IBookDo, queryFilter *dto.BookListQueryFilter) query.IBookDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.Book.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.Book.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.Title != nil {
		query = query.Where(dao.query.Book.Title.Eq(*queryFilter.Title))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.Book.Description.Eq(*queryFilter.Description))
	}
	if queryFilter.Content != nil {
		query = query.Where(dao.query.Book.Content.Eq(*queryFilter.Content))
	}
	if queryFilter.Author != nil {
		query = query.Where(dao.query.Book.Author.Eq(*queryFilter.Author))
	}
	if queryFilter.Source != nil {
		query = query.Where(dao.query.Book.Source.Eq(*queryFilter.Source))
	}
	if queryFilter.Isbn != nil {
		query = query.Where(dao.query.Book.Isbn.Eq(*queryFilter.Isbn))
	}
	if queryFilter.Price != nil {
		query = query.Where(dao.query.Book.Price.Eq(*queryFilter.Price))
	}

	return query
}

func (dao *BookDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Book.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *BookDao) Update(ctx context.Context, model *models.Book) error {
	_, err := dao.Context(ctx).Where(dao.query.Book.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *BookDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Book.ID.Eq(id)).Delete()
	return err
}

func (dao *BookDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Book.ID.Eq(id)).Delete()
	return err
}

func (dao *BookDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Book.ID.Eq(id)).UpdateSimple(dao.query.Book.DeletedAt.Null())
	return err
}

func (dao *BookDao) Create(ctx context.Context, model *models.Book) error {
	return dao.Context(ctx).Create(model)
}

func (dao *BookDao) GetByID(ctx context.Context, id uint64) (*models.Book, error) {
	return dao.Context(ctx).Where(dao.query.Book.ID.Eq(id)).First()
}

func (dao *BookDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Book, error) {
	return dao.Context(ctx).Where(dao.query.Book.ID.In(ids...)).Find()
}

func (dao *BookDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.BookListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Book, int64, error) {
	query := dao.query.Book
	bookQuery := query.WithContext(ctx)
	bookQuery = dao.decorateQueryFilter(bookQuery, queryFilter)
	bookQuery = dao.decorateSortQueryFilter(bookQuery, sortFilter)
	return bookQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *BookDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.BookListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Book, error) {
	query := dao.query.Book
	bookQuery := query.WithContext(ctx)
	bookQuery = dao.decorateQueryFilter(bookQuery, queryFilter)
	bookQuery = dao.decorateSortQueryFilter(bookQuery, sortFilter)
	return bookQuery.Find()
}

func (dao *BookDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.BookListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.Book, error) {
	query := dao.query.Book
	bookQuery := query.WithContext(ctx)
	bookQuery = dao.decorateQueryFilter(bookQuery, queryFilter)
	bookQuery = dao.decorateSortQueryFilter(bookQuery, sortFilter)
	return bookQuery.First()
}
