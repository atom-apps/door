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
type ChapterDao struct {
	query *query.Query
}

func (dao *ChapterDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *ChapterDao) Context(ctx context.Context) query.IChapterDo {
	return dao.query.Chapter.WithContext(ctx)
}

func (dao *ChapterDao) decorateSortQueryFilter(query query.IChapterDo, sortFilter *ds.SortQueryFilter) query.IChapterDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Chapter.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Chapter.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *ChapterDao) decorateQueryFilter(query query.IChapterDo, queryFilter *dto.ChapterListQueryFilter) query.IChapterDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.Chapter.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.Chapter.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.BookID != nil {
		query = query.Where(dao.query.Chapter.BookID.Eq(*queryFilter.BookID))
	}
	if queryFilter.Title != nil {
		query = query.Where(dao.query.Chapter.Title.Eq(*queryFilter.Title))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.Chapter.Description.Eq(*queryFilter.Description))
	}
	if queryFilter.Content != nil {
		query = query.Where(dao.query.Chapter.Content.Eq(*queryFilter.Content))
	}

	return query
}

func (dao *ChapterDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Chapter.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ChapterDao) Update(ctx context.Context, model *models.Chapter) error {
	_, err := dao.Context(ctx).Where(dao.query.Chapter.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ChapterDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Chapter.ID.Eq(id)).Delete()
	return err
}

func (dao *ChapterDao) DeleteByBookID(ctx context.Context, bookID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Chapter.BookID.Eq(bookID)).Delete()
	return err
}

func (dao *ChapterDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Chapter.ID.Eq(id)).Delete()
	return err
}

func (dao *ChapterDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Chapter.ID.Eq(id)).UpdateSimple(dao.query.Chapter.DeletedAt.Null())
	return err
}

func (dao *ChapterDao) Create(ctx context.Context, model *models.Chapter) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ChapterDao) GetByID(ctx context.Context, id uint64) (*models.Chapter, error) {
	return dao.Context(ctx).Where(dao.query.Chapter.ID.Eq(id)).First()
}

func (dao *ChapterDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Chapter, error) {
	return dao.Context(ctx).Where(dao.query.Chapter.ID.In(ids...)).Find()
}

func (dao *ChapterDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ChapterListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Chapter, int64, error) {
	query := dao.query.Chapter
	chapterQuery := query.WithContext(ctx)
	chapterQuery = dao.decorateQueryFilter(chapterQuery, queryFilter)
	chapterQuery = dao.decorateSortQueryFilter(chapterQuery, sortFilter)
	return chapterQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *ChapterDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ChapterListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Chapter, error) {
	query := dao.query.Chapter
	chapterQuery := query.WithContext(ctx)
	chapterQuery = dao.decorateQueryFilter(chapterQuery, queryFilter)
	chapterQuery = dao.decorateSortQueryFilter(chapterQuery, sortFilter)
	return chapterQuery.Find()
}

func (dao *ChapterDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ChapterListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.Chapter, error) {
	query := dao.query.Chapter
	chapterQuery := query.WithContext(ctx)
	chapterQuery = dao.decorateQueryFilter(chapterQuery, queryFilter)
	chapterQuery = dao.decorateSortQueryFilter(chapterQuery, sortFilter)
	return chapterQuery.First()
}

// CountByBookID
func (dao *ChapterDao) CountByBookID(ctx context.Context, bookID uint64) (int64, error) {
	return dao.Context(ctx).Where(dao.query.Chapter.BookID.Eq(bookID)).Count()
}
