package dao

import (
	"context"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"gorm.io/gen/field"
)

// @provider
type ArticleForwardSourceDao struct {
	query *query.Query
}

func (dao *ArticleForwardSourceDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *ArticleForwardSourceDao) Context(ctx context.Context) query.IArticleForwardSourceDo {
	return dao.query.ArticleForwardSource.WithContext(ctx)
}

func (dao *ArticleForwardSourceDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleForwardSource.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleForwardSourceDao) Update(ctx context.Context, model *models.ArticleForwardSource) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleForwardSource.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleForwardSourceDao) DeleteByArticleID(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleForwardSource.ArticleID.Eq(id)).Delete()
	return err
}

func (dao *ArticleForwardSourceDao) Create(ctx context.Context, model *models.ArticleForwardSource) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleForwardSourceDao) GetByID(ctx context.Context, id uint64) (*models.ArticleForwardSource, error) {
	return dao.Context(ctx).Where(dao.query.ArticleForwardSource.ID.Eq(id)).First()
}

func (dao *ArticleForwardSourceDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleForwardSource, error) {
	return dao.Context(ctx).Where(dao.query.ArticleForwardSource.ID.In(ids...)).Find()
}
