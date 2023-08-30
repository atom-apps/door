package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/test/dao"
	"github.com/atom-apps/door/modules/test/dto"

	"github.com/jinzhu/copier"
)

// @provider
type UserService struct {
	userDao *dao.UserDao
}

func (svc *UserService) DecorateItem(model *models.User, id int) *dto.UserItem {
	var dtoItem *dto.UserItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return svc.userDao.GetByID(ctx, id)
}

func (svc *UserService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, error) {
	return svc.userDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *UserService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, int64, error) {
	return svc.userDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *UserService) CreateFromModel(ctx context.Context, model *models.User) error {
	return svc.userDao.Create(ctx, model)
}

// Create
func (svc *UserService) Create(ctx context.Context, body *dto.UserForm) error {
	model := &models.User{}
	_ = copier.Copy(model, body)
	return svc.userDao.Create(ctx, model)
}

// Update
func (svc *UserService) Update(ctx context.Context, id int64, body *dto.UserForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.userDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *UserService) UpdateFromModel(ctx context.Context, model *models.User) error {
	return svc.userDao.Update(ctx, model)
}

// Delete
func (svc *UserService) Delete(ctx context.Context, id int64) error {
	return svc.userDao.Delete(ctx, id)
}
