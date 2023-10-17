package dao

import (
	"context"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"gorm.io/gen/field"
)

// @provider
type ArticlePaidUserDao struct {
	query *query.Query
}

func (dao *ArticlePaidUserDao) Context(ctx context.Context) query.IArticlePaidUserDo {
	return dao.query.ArticlePaidUser.WithContext(ctx)
}

func (dao *ArticlePaidUserDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticlePaidUser.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticlePaidUserDao) Update(ctx context.Context, model *models.ArticlePaidUser) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticlePaidUser.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticlePaidUserDao) DeleteByArticleID(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticlePaidUser.ArticleID.Eq(id)).Delete()
	return err
}

func (dao *ArticlePaidUserDao) Create(ctx context.Context, model *models.ArticlePaidUser) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticlePaidUserDao) GetByID(ctx context.Context, id uint64) (*models.ArticlePaidUser, error) {
	return dao.Context(ctx).Where(dao.query.ArticlePaidUser.ID.Eq(id)).First()
}

func (dao *ArticlePaidUserDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticlePaidUser, error) {
	return dao.Context(ctx).Where(dao.query.ArticlePaidUser.ID.In(ids...)).Find()
}
