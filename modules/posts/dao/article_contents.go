package dao

import (
	"context"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"gorm.io/gen/field"
)

// @provider
type ArticleContentDao struct {
	query *query.Query
}

func (dao *ArticleContentDao) Context(ctx context.Context) query.IArticleContentDo {
	return dao.query.ArticleContent.WithContext(ctx)
}

func (dao *ArticleContentDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleContentDao) Update(ctx context.Context, model *models.ArticleContent) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleContentDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleContentDao) Create(ctx context.Context, model *models.ArticleContent) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleContentDao) GetByID(ctx context.Context, id uint64) (*models.ArticleContent, error) {
	return dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(id)).First()
}

func (dao *ArticleContentDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleContent, error) {
	return dao.Context(ctx).Where(dao.query.ArticleContent.ID.In(ids...)).Find()
}
