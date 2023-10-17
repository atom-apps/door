package dao

import (
	"context"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"gorm.io/gen/field"
)

// @provider
type ArticleDigDao struct {
	query *query.Query
}

func (dao *ArticleDigDao) Context(ctx context.Context) query.IArticleDigDo {
	return dao.query.ArticleDig.WithContext(ctx)
}

func (dao *ArticleDigDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleDig.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleDigDao) Update(ctx context.Context, model *models.ArticleDig) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleDig.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleDigDao) DeleteByArticleID(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleDig.ArticleID.Eq(id)).Delete()
	return err
}

func (dao *ArticleDigDao) Create(ctx context.Context, model *models.ArticleDig) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleDigDao) GetByID(ctx context.Context, id uint64) (*models.ArticleDig, error) {
	return dao.Context(ctx).Where(dao.query.ArticleDig.ID.Eq(id)).First()
}

func (dao *ArticleDigDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleDig, error) {
	return dao.Context(ctx).Where(dao.query.ArticleDig.ID.In(ids...)).Find()
}
