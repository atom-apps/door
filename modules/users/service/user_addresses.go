package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"

	"github.com/jinzhu/copier"
)

// @provider
type UserAddressService struct {
	userAddressDao *dao.UserAddressDao
}

func (svc *UserAddressService) DecorateItem(model *models.UserAddress, id int) *dto.UserAddressItem {
	return &dto.UserAddressItem{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		UserID:    model.UserID,
		Code:      model.Code,
		Town:      model.Town,
		Detail:    model.Detail,
		Name:      model.Name,
		Phone:     model.Phone,
		ZipCode:   model.ZipCode,
	}
}

func (svc *UserAddressService) GetByID(ctx context.Context, id uint64) (*models.UserAddress, error) {
	return svc.userAddressDao.GetByID(ctx, id)
}

func (svc *UserAddressService) FindByQueryFilter(
	ctx context.Context,
	userId uint64,
	queryFilter *dto.UserAddressListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserAddress, error) {
	return svc.userAddressDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *UserAddressService) PageByQueryFilter(
	ctx context.Context,
	userId uint64,
	queryFilter *dto.UserAddressListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserAddress, int64, error) {
	return svc.userAddressDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *UserAddressService) CreateFromModel(ctx context.Context, model *models.UserAddress) error {
	return svc.userAddressDao.Create(ctx, model)
}

// Create
func (svc *UserAddressService) Create(ctx context.Context, userId uint64, body *dto.UserAddressForm) error {
	model := &models.UserAddress{}
	_ = copier.Copy(model, body)
	model.UserID = userId
	return svc.userAddressDao.Create(ctx, model)
}

// Update
func (svc *UserAddressService) Update(ctx context.Context, userId uint64, id uint64, body *dto.UserAddressForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	model.UserID = userId
	return svc.userAddressDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *UserAddressService) UpdateFromModel(ctx context.Context, model *models.UserAddress) error {
	return svc.userAddressDao.Update(ctx, model)
}

// Delete
func (svc *UserAddressService) Delete(ctx context.Context, userId uint64, id uint64) error {
	return svc.userAddressDao.Delete(ctx, id)
}

// FindByUserID
func (svc *UserAddressService) FindByUserID(ctx context.Context, userID uint64) ([]*models.UserAddress, error) {
	return svc.userAddressDao.FindByUserID(ctx, userID)
}
