package dao

import (
	"context"
	"time"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/posts/dto"
	"github.com/samber/lo"
	"gorm.io/gen/field"
)

// @provider
type ArticleDao struct {
	query  *query.Query
	tagDao *ArticleTagDao
}

func (dao *ArticleDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *ArticleDao) Context(ctx context.Context) query.IArticleDo {
	return dao.query.Article.WithContext(ctx)
}

func (dao *ArticleDao) decorateSortQueryFilter(query query.IArticleDo, sortFilter *ds.SortQueryFilter) query.IArticleDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Article.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Article.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	orderExprs = append(orderExprs, dao.query.Article.Weight.Desc())
	return query.Order(orderExprs...)
}

func (dao *ArticleDao) decorateQueryFilter(query query.IArticleDo, queryFilter *dto.ArticleListQueryFilter) query.IArticleDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.Article.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.Article.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.UUID != nil {
		query = query.Where(dao.query.Article.UUID.Eq(*queryFilter.UUID))
	}
	if queryFilter.BookID != nil {
		query = query.Where(dao.query.Article.BookID.Eq(*queryFilter.BookID))
	}
	if queryFilter.ChapterID != nil {
		query = query.Where(dao.query.Article.ChapterID.Eq(*queryFilter.ChapterID))
	}
	if queryFilter.CategoryID != nil {
		query = query.Where(dao.query.Article.CategoryID.Eq(*queryFilter.CategoryID))
	}
	if queryFilter.Published != nil {
		query = query.Where(query.Where(dao.query.Article.PublishAt.Lt(time.Now())).Or(dao.query.Article.PublishAt.IsNull()))
	}
	if queryFilter.Type != nil {
		query = query.Where(dao.query.Article.Type.Eq(*queryFilter.Type))
	}
	if queryFilter.Format != nil {
		query = query.Where(dao.query.Article.Format.Eq(*queryFilter.Format))
	}
	if queryFilter.Title != nil {
		query = query.Where(dao.query.Article.Title.Eq(*queryFilter.Title))
	}
	if queryFilter.Keyword != nil {
		query = query.Where(dao.query.Article.Keyword.Eq(*queryFilter.Keyword))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.Article.Description.Eq(*queryFilter.Description))
	}
	if queryFilter.PostIP != nil {
		query = query.Where(dao.query.Article.PostIP.Eq(*queryFilter.PostIP))
	}

	return query
}

func (dao *ArticleDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Article.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleDao) Update(ctx context.Context, model *models.Article) error {
	_, err := dao.Context(ctx).Where(dao.query.Article.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Article.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleDao) DeleteByBookID(ctx context.Context, bookID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Article.BookID.Eq(bookID)).Delete()
	return err
}

func (dao *ArticleDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Article.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Article.ID.Eq(id)).UpdateSimple(dao.query.Article.DeletedAt.Null())
	return err
}

func (dao *ArticleDao) Create(ctx context.Context, model *models.Article) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleDao) GetByID(ctx context.Context, id uint64) (*dto.ArticleItem, error) {
	type articleQueryItem struct {
		models.Article
		// content
		FreeContent  string
		PriceContent string
		// digs
		Views    uint64
		Likes    uint64
		Dislikes uint64
		// source
		Source       string
		SourceAuthor string
	}

	article, content, dig, forward := dao.query.Article, dao.query.ArticleContent, dao.query.ArticleDig, dao.query.ArticleForwardSource

	var item *articleQueryItem

	fields := []field.Expr{
		article.ALL,
		// content
		content.FreeContent,
		content.PriceContent,

		// dig
		dig.Views,
		dig.Likes,
		dig.Dislikes,
		// forwards
		forward.Source,
		forward.SourceAuthor,
	}
	err := dao.Context(ctx).
		Select(fields...).
		LeftJoin(dig, dig.ArticleID.EqCol(article.ID)).
		LeftJoin(forward, forward.ArticleID.EqCol(article.ID)).
		LeftJoin(content, content.ArticleID.EqCol(article.ID)).
		Scan(&item)
	if err != nil {
		return nil, err
	}

	attachmentModels, err := dao.query.ArticleAttachment.WithContext(ctx).Where(dao.query.ArticleAttachment.ArticleID.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	attachments := lo.GroupBy(attachmentModels, func(item *models.ArticleAttachment) uint64 { return item.ID })

	paymentModels, err := dao.query.ArticlePayment.WithContext(ctx).Where(dao.query.ArticlePayment.ArticleID.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	payments := lo.GroupBy(paymentModels, func(item *models.ArticlePayment) uint64 { return item.ID })

	ret := &dto.ArticleItem{
		FormatCN: item.Format.Cn(),
		TypeCN:   item.Type.Cn(),
	}

	if err := ret.Fill(item); err != nil {
		return nil, err
	}

	ret.Dig = dto.ArticleDigItemFillWith(item)
	ret.ForwardSource = dto.ArticleForwardSourceItemFillWith(item)
	ret.Content = &dto.ArticleContentItem{
		FreeContent:  item.FreeContent,
		PriceContent: item.PriceContent,
	}

	if tags, err := dao.tagDao.GetByArticleID(ctx, item.ID); err == nil {
		ret.Tags = tags
	}

	if resources, ok := attachments[item.ID]; ok {
		ret.Attachments = lo.Map(resources, func(item *models.ArticleAttachment, _ int) *dto.ArticleAttachmentItem {
			return &dto.ArticleAttachmentItem{
				FilesystemID: item.FilesystemID,
				Description:  item.Description,
				Password:     item.Password,
			}
		})
	}

	// payments
	if resources, ok := payments[item.ID]; ok {
		ret.Payments = lo.Map(resources, func(item *models.ArticlePayment, _ int) *dto.ArticlePaymentItem {
			return &dto.ArticlePaymentItem{
				PriceType:       item.PriceType,
				Token:           item.Token,
				Price:           item.Price,
				Discount:        item.Discount,
				DiscountStartAt: item.DiscountStartAt,
				DiscountEndAt:   item.DiscountEndAt,
			}
		})
	}

	return ret, err
}

func (dao *ArticleDao) GetModelByID(ctx context.Context, id uint64) (*models.Article, error) {
	return dao.Context(ctx).Where(dao.query.Article.ID.Eq(id)).First()
}

func (dao *ArticleDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Article, error) {
	return dao.Context(ctx).Where(dao.query.Article.ID.In(ids...)).Find()
}

func (dao *ArticleDao) PageByQueryFilter(ctx context.Context, queryFilter *dto.ArticleListQueryFilter, pageFilter *ds.PageQueryFilter, sortFilter *ds.SortQueryFilter) ([]*dto.ArticleItem, int64, error) {
	type articleQueryItem struct {
		models.Article
		// digs
		Views    uint64
		Likes    uint64
		Dislikes uint64
		// source
		Source       string
		SourceAuthor string
	}

	article, dig, forward := dao.query.Article, dao.query.ArticleDig, dao.query.ArticleForwardSource

	query := dao.Context(ctx)
	query = dao.decorateQueryFilter(query, queryFilter)
	query = dao.decorateSortQueryFilter(query, sortFilter)

	var items []*articleQueryItem

	fields := []field.Expr{
		article.ALL,
		// dig
		dig.Views,
		dig.Likes,
		dig.Dislikes,
		// forwards
		forward.Source,
		forward.SourceAuthor,
	}

	count, err := query.
		Select(fields...).
		LeftJoin(dig, dig.ArticleID.EqCol(article.ID)).
		LeftJoin(forward, forward.ArticleID.EqCol(article.ID)).
		ScanByPage(&items, pageFilter.Offset(), pageFilter.Limit)
	if err != nil {
		return nil, 0, err
	}

	ids := lo.Map(items, func(item *articleQueryItem, _ int) uint64 {
		return item.ID
	})

	attachmentModels, err := dao.query.ArticleAttachment.WithContext(ctx).Where(dao.query.ArticleAttachment.ArticleID.In(ids...)).Find()
	if err != nil {
		return nil, 0, err
	}
	attachments := lo.GroupBy(attachmentModels, func(item *models.ArticleAttachment) uint64 { return item.ID })

	paymentModels, err := dao.query.ArticlePayment.WithContext(ctx).Where(dao.query.ArticlePayment.ArticleID.In(ids...)).Find()
	if err != nil {
		return nil, 0, err
	}
	payments := lo.GroupBy(paymentModels, func(item *models.ArticlePayment) uint64 { return item.ID })

	tags, err := dao.tagDao.GetByArticleIDs(ctx, ids)
	if err != nil {
		return nil, 0, err
	}

	results := lo.Map(items, func(item *articleQueryItem, _ int) *dto.ArticleItem {
		ret := &dto.ArticleItem{
			FormatCN: item.Format.Cn(),
			TypeCN:   item.Type.Cn(),
		}

		ret.Fill(item)
		ret.Dig = dto.ArticleDigItemFillWith(item)
		ret.ForwardSource = dto.ArticleForwardSourceItemFillWith(item)
		ret.Content = &dto.ArticleContentItem{
			FreeContent:  "",
			PriceContent: "",
		}

		if resources, ok := tags[item.ID]; ok {
			ret.Tags = resources
		}

		if resources, ok := attachments[item.ID]; ok {
			ret.Attachments = lo.Map(resources, func(item *models.ArticleAttachment, _ int) *dto.ArticleAttachmentItem {
				return &dto.ArticleAttachmentItem{
					FilesystemID: item.FilesystemID,
					Description:  item.Description,
					Password:     item.Password,
				}
			})
		}

		// payments
		if resources, ok := payments[item.ID]; ok {
			ret.Payments = lo.Map(resources, func(item *models.ArticlePayment, _ int) *dto.ArticlePaymentItem {
				return &dto.ArticlePaymentItem{
					PriceType:       item.PriceType,
					Token:           item.Token,
					Price:           item.Price,
					Discount:        item.Discount,
					DiscountStartAt: item.DiscountStartAt,
					DiscountEndAt:   item.DiscountEndAt,
				}
			})
		}

		return ret
	})

	return results, count, err
}

// CountByBookID
func (dao *ArticleDao) CountByBookID(ctx context.Context, bookID uint64) (int64, error) {
	return dao.Context(ctx).Where(dao.query.Article.BookID.Eq(bookID)).Count()
}

// CountByChapterID
func (dao *ArticleDao) CountByChapterID(ctx context.Context, chapterID uint64) (int64, error) {
	return dao.Context(ctx).Where(dao.query.Article.ChapterID.Eq(chapterID)).Count()
}
