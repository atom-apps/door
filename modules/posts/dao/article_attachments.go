package dao

import (
	"context"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"gorm.io/gen/field"
)

// @provider
type ArticleAttachmentDao struct {
	query *query.Query
}

func (dao *ArticleAttachmentDao) Context(ctx context.Context) query.IArticleAttachmentDo {
	return dao.query.ArticleAttachment.WithContext(ctx)
}

func (dao *ArticleAttachmentDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleAttachmentDao) Update(ctx context.Context, model *models.ArticleAttachment) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleAttachmentDao) DeleteByArticleID(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleAttachment.ArticleID.Eq(id)).Delete()
	return err
}

func (dao *ArticleAttachmentDao) Create(ctx context.Context, model *models.ArticleAttachment) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleAttachmentDao) GetByID(ctx context.Context, id uint64) (*models.ArticleAttachment, error) {
	return dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.Eq(id)).First()
}

func (dao *ArticleAttachmentDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleAttachment, error) {
	return dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.In(ids...)).Find()
}
