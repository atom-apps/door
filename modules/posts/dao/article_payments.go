package dao

import (
	"context"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"gorm.io/gen/field"
)

// @provider
type ArticlePaymentDao struct {
	query *query.Query
}

func (dao *ArticlePaymentDao) Context(ctx context.Context) query.IArticlePaymentDo {
	return dao.query.ArticlePayment.WithContext(ctx)
}

func (dao *ArticlePaymentDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticlePayment.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticlePaymentDao) Update(ctx context.Context, model *models.ArticlePayment) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticlePayment.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticlePaymentDao) DeleteByArticleID(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticlePayment.ArticleID.Eq(id)).Delete()
	return err
}

func (dao *ArticlePaymentDao) Create(ctx context.Context, model *models.ArticlePayment) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticlePaymentDao) GetByID(ctx context.Context, id uint64) (*models.ArticlePayment, error) {
	return dao.Context(ctx).Where(dao.query.ArticlePayment.ID.Eq(id)).First()
}

func (dao *ArticlePaymentDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticlePayment, error) {
	return dao.Context(ctx).Where(dao.query.ArticlePayment.ID.In(ids...)).Find()
}
