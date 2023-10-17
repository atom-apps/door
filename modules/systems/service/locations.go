package service

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/systems/dto"

	"github.com/jinzhu/copier"
)

// @provider
type LocationService struct {
	locationDao *dao.LocationDao
}

func (svc *LocationService) DecorateItem(model *models.Location, id int) *dto.LocationItem {
	return &dto.LocationItem{
		ID:       model.ID,
		Code:     model.Code,
		Name:     model.Name,
		Province: model.Province,
		City:     model.City,
		Area:     model.Area,
		Town:     model.Town,
	}
}

func (svc *LocationService) GetByID(ctx context.Context, id uint64) (*models.Location, error) {
	return svc.locationDao.GetByID(ctx, id)
}

func (svc *LocationService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.LocationListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Location, error) {
	return svc.locationDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *LocationService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.LocationListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
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
func (svc *LocationService) Update(ctx context.Context, id uint64, body *dto.LocationForm) error {
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
func (svc *LocationService) Delete(ctx context.Context, id uint64) error {
	return svc.locationDao.Delete(ctx, id)
}

// Provinces
func (svc *LocationService) Provinces(ctx context.Context) ([]*models.Location, error) {
	return svc.locationDao.Provinces(ctx)
}

// Cities
func (svc *LocationService) Cities(ctx context.Context, province string) ([]*models.Location, error) {
	return svc.locationDao.Cities(ctx, province)
}

// Area
func (svc *LocationService) Areas(ctx context.Context, province, city string) ([]*models.Location, error) {
	return svc.locationDao.Areas(ctx, province, city)
}

// Towns
func (svc *LocationService) Towns(ctx context.Context, province, city, area string) ([]*models.Location, error) {
	return svc.locationDao.Towns(ctx, province, city, area)
}

// GetByCode
func (svc *LocationService) GetByCode(ctx context.Context, code string) (*models.Location, error) {
	return svc.locationDao.GetByCode(ctx, code)
}

func (svc *LocationService) ParseCode(ctx context.Context, code string) (string, string, string) {
	return code[:2], code[2:2], code[4:]
}

// GetByCode
func (svc *LocationService) GetByCodeTown(ctx context.Context, code, townCode string) (*dto.LocationDetail, error) {
	provinceCode, cityCode, areaCode := svc.ParseCode(ctx, code)

	province, err := svc.locationDao.Province(ctx, provinceCode)
	if err != nil {
		return nil, err
	}

	city, err := svc.locationDao.City(ctx, provinceCode, cityCode)
	if err != nil {
		return nil, err
	}

	area, err := svc.locationDao.Area(ctx, provinceCode, cityCode, areaCode)
	if err != nil {
		return nil, err
	}

	town, err := svc.locationDao.Town(ctx, provinceCode, cityCode, areaCode, townCode)
	if err != nil {
		return nil, err
	}

	return &dto.LocationDetail{
		Code:     code,
		TownCode: townCode,
		Province: province.Name,
		City:     city.Name,
		Area:     area.Name,
		Town:     town.Name,
	}, nil
}
