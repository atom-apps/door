package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/tools/dao"
	"github.com/atom-apps/door/modules/tools/dto"

	"github.com/jinzhu/copier"
)

// @provider
type LocationService struct {
	locationDao *dao.LocationDao
}

func (svc *LocationService) DecorateItem(model *models.Location, id int) *dto.LocationItem {
	var dtoItem *dto.LocationItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *LocationService) GetByID(ctx context.Context, id int64) (*models.Location, error) {
	return svc.locationDao.GetByID(ctx, id)
}

func (svc *LocationService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.LocationListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Location, error) {
	return svc.locationDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *LocationService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.LocationListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Location, int64, error) {
	return svc.locationDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *LocationService) CreateFromModel(ctx context.Context, model *models.Location) error {
	return svc.locationDao.Create(ctx, model)
}

// Create
func (svc *LocationService) Create(ctx context.Context, body *dto.LocationForm) error {
	model := &models.Location{}
	_ = copier.Copy(model, body)
	return svc.locationDao.Create(ctx, model)
}

// Update
func (svc *LocationService) Update(ctx context.Context, id int64, body *dto.LocationForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.locationDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *LocationService) UpdateFromModel(ctx context.Context, model *models.Location) error {
	return svc.locationDao.Update(ctx, model)
}

// Delete
func (svc *LocationService) Delete(ctx context.Context, id int64) error {
	return svc.locationDao.Delete(ctx, id)
}
