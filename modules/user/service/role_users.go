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
type RoleUserService struct {
	roleUserDao *dao.RoleUserDao
}

func (svc *RoleUserService) DecorateItem(model *models.RoleUser, id int) *dto.RoleUserItem {
	var dtoItem *dto.RoleUserItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *RoleUserService) GetByID(ctx context.Context, id int64) (*models.RoleUser, error) {
	return svc.roleUserDao.GetByID(ctx, id)
}

func (svc *RoleUserService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleUserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.RoleUser, error) {
	return svc.roleUserDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *RoleUserService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleUserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.RoleUser, int64, error) {
	return svc.roleUserDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *RoleUserService) CreateFromModel(ctx context.Context, model *models.RoleUser) error {
	return svc.roleUserDao.Create(ctx, model)
}

// Create
func (svc *RoleUserService) Create(ctx context.Context, body *dto.RoleUserForm) error {
	model := &models.RoleUser{}
	_ = copier.Copy(model, body)
	return svc.roleUserDao.Create(ctx, model)
}

// Update
func (svc *RoleUserService) Update(ctx context.Context, id int64, body *dto.RoleUserForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.roleUserDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *RoleUserService) UpdateFromModel(ctx context.Context, model *models.RoleUser) error {
	return svc.roleUserDao.Update(ctx, model)
}

// Delete
func (svc *RoleUserService) Delete(ctx context.Context, id int64) error {
	return svc.roleUserDao.Delete(ctx, id)
}
