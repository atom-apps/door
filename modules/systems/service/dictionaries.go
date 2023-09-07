package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/systems/dto"

	"github.com/jinzhu/copier"
)

// @provider
type DictionaryService struct {
	dictionaryDao *dao.DictionaryDao
}

func (svc *DictionaryService) DecorateItem(model *models.Dictionary, id int) *dto.DictionaryItem {
	return &dto.DictionaryItem{
		ID:          model.ID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		Items:       model.Items,
	}
}

func (svc *DictionaryService) GetByID(ctx context.Context, id uint64) (*models.Dictionary, error) {
	return svc.dictionaryDao.GetByID(ctx, id)
}

func (svc *DictionaryService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Dictionary, error) {
	return svc.dictionaryDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *DictionaryService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DictionaryListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Dictionary, int64, error) {
	return svc.dictionaryDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *DictionaryService) CreateFromModel(ctx context.Context, model *models.Dictionary) error {
	return svc.dictionaryDao.Create(ctx, model)
}

// Create
func (svc *DictionaryService) Create(ctx context.Context, body *dto.DictionaryForm) error {
	model := &models.Dictionary{}
	_ = copier.Copy(model, body)
	return svc.dictionaryDao.Create(ctx, model)
}

// Update
func (svc *DictionaryService) Update(ctx context.Context, id uint64, body *dto.DictionaryForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.dictionaryDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *DictionaryService) UpdateFromModel(ctx context.Context, model *models.Dictionary) error {
	return svc.dictionaryDao.Update(ctx, model)
}

// Delete
func (svc *DictionaryService) Delete(ctx context.Context, id uint64) error {
	return svc.dictionaryDao.Delete(ctx, id)
}
