package service

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/posts/dao"
	"github.com/atom-apps/door/modules/posts/dto"

	"github.com/jinzhu/copier"
)

// @provider
type ArticleService struct {
	articleDao       *dao.ArticleDao
	attachmentDao    *dao.ArticleAttachmentDao
	contentDao       *dao.ArticleContentDao
	digDao           *dao.ArticleDigDao
	forwardSourceDao *dao.ArticleForwardSourceDao
	paidUserDao      *dao.ArticlePaidUserDao
	paymentDao       *dao.ArticlePaymentDao
}

func (svc *ArticleService) DecorateItem(item *dto.ArticleItem, id int) *dto.ArticleItem {
	return item
}

func (svc *ArticleService) GetByID(ctx context.Context, id uint64) (*dto.ArticleItem, error) {
	return svc.articleDao.GetByID(ctx, id)
}

func (svc *ArticleService) PageByQueryFilter(ctx context.Context, queryFilter *dto.ArticleListQueryFilter, pageFilter *ds.PageQueryFilter, sortFilter *ds.SortQueryFilter) ([]*dto.ArticleItem, int64, error) {
	return svc.articleDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *ArticleService) CreateFromModel(ctx context.Context, model *models.Article) error {
	return svc.articleDao.Create(ctx, model)
}

// Create
func (svc *ArticleService) Create(ctx context.Context, body *dto.ArticleForm) error {
	model := &models.Article{}
	_ = copier.Copy(model, body)
	return svc.articleDao.Create(ctx, model)
}

// Update
func (svc *ArticleService) Update(ctx context.Context, id uint64, body *dto.ArticleForm) error {
	model, err := svc.articleDao.GetModelByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.articleDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *ArticleService) UpdateFromModel(ctx context.Context, model *models.Article) error {
	return svc.articleDao.Update(ctx, model)
}

// Delete
func (svc *ArticleService) Delete(ctx context.Context, id uint64) error {
	return svc.articleDao.Delete(ctx, id)
}
