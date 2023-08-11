package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/user/dao"
	"github.com/atom-apps/door/modules/user/dto"

	"github.com/jinzhu/copier"
)

// @provider
type UserInfoService struct {
	userInfoDao *dao.UserInfoDao
}

func (svc *UserInfoService) DecorateItem(model *models.UserInfo, id int) *dto.UserInfoItem {
	var dtoItem *dto.UserInfoItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *UserInfoService) GetByID(ctx context.Context, id int64) (*models.UserInfo, error) {
	return svc.userInfoDao.GetByID(ctx, id)
}

func (svc *UserInfoService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserInfoListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserInfo, error) {
	return svc.userInfoDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *UserInfoService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserInfoListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserInfo, int64, error) {
	return svc.userInfoDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *UserInfoService) CreateFromModel(ctx context.Context, model *models.UserInfo) error {
	return svc.userInfoDao.Create(ctx, model)
}

// Create
func (svc *UserInfoService) Create(ctx context.Context, body *dto.UserInfoForm) error {
	model := &models.UserInfo{}
	_ = copier.Copy(model, body)
	return svc.userInfoDao.Create(ctx, model)
}

// Update
func (svc *UserInfoService) Update(ctx context.Context, id int64, body *dto.UserInfoForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.userInfoDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *UserInfoService) UpdateFromModel(ctx context.Context, model *models.UserInfo) error {
	return svc.userInfoDao.Update(ctx, model)
}

// Delete
func (svc *UserInfoService) Delete(ctx context.Context, id int64) error {
	return svc.userInfoDao.Delete(ctx, id)
}
