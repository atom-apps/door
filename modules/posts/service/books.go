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
type BookService struct {
	bookDao    *dao.BookDao
	chapterDao *dao.ChapterDao
	articleDao *dao.ArticleDao
}

func (svc *BookService) DecorateItem(model *models.Book, id int) *dto.BookItem {
	item := &dto.BookItem{
		ID:           model.ID,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		TenantID:     model.TenantID,
		UserID:       model.UserID,
		Title:        model.Title,
		Description:  model.Description,
		Content:      model.Content,
		Author:       model.Author,
		Source:       model.Source,
		Isbn:         model.Isbn,
		Price:        model.Price,
		ChapterCount: 0,
		ArticleCount: 0,
	}

	if count, err := svc.chapterDao.CountByBookID(context.Background(), model.ID); err == nil {
		item.ChapterCount = count
	}

	if count, err := svc.articleDao.CountByBookID(context.Background(), model.ID); err == nil {
		item.ArticleCount = count
	}
	return item
}

func (svc *BookService) GetByID(ctx context.Context, id uint64) (*models.Book, error) {
	return svc.bookDao.GetByID(ctx, id)
}

func (svc *BookService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.BookListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Book, error) {
	return svc.bookDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *BookService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.BookListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Book, int64, error) {
	return svc.bookDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *BookService) CreateFromModel(ctx context.Context, model *models.Book) error {
	return svc.bookDao.Create(ctx, model)
}

// Create
func (svc *BookService) Create(ctx context.Context, body *dto.BookForm) error {
	model := &models.Book{}
	_ = copier.Copy(model, body)
	return svc.bookDao.Create(ctx, model)
}

// Update
func (svc *BookService) Update(ctx context.Context, id uint64, body *dto.BookForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.bookDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *BookService) UpdateFromModel(ctx context.Context, model *models.Book) error {
	return svc.bookDao.Update(ctx, model)
}

// Delete
func (svc *BookService) Delete(ctx context.Context, id uint64) error {
	return svc.bookDao.Transaction(func() error {
		if err := svc.chapterDao.DeleteByBookID(ctx, id); err != nil {
			return err
		}
		if err := svc.articleDao.DeleteByBookID(ctx, id); err != nil {
			return err
		}

		return svc.bookDao.Delete(ctx, id)
	})
}
