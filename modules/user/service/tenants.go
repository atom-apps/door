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
type TenantService struct {
	tenantDao         *dao.TenantDao
	permissionRuleSvc *PermissionRuleService
}

func (svc *TenantService) DecorateItem(model *models.Tenant, id int) *dto.TenantItem {
	return &dto.TenantItem{
		Name:        model.Name,
		Description: model.Description,
		Meta:        model.Meta,
	}
}

func (svc *TenantService) GetByID(ctx context.Context, id int64) (*models.Tenant, error) {
	return svc.tenantDao.GetByID(ctx, id)
}

func (svc *TenantService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Tenant, error) {
	return svc.tenantDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *TenantService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Tenant, int64, error) {
	return svc.tenantDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *TenantService) CreateFromModel(ctx context.Context, model *models.Tenant) error {
	return svc.tenantDao.Create(ctx, model)
}

// Create
func (svc *TenantService) Create(ctx context.Context, body *dto.TenantForm) error {
	model := &models.Tenant{}
	_ = copier.Copy(model, body)
	return svc.tenantDao.Create(ctx, model)
}

// Update
func (svc *TenantService) Update(ctx context.Context, id int64, body *dto.TenantForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.tenantDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *TenantService) UpdateFromModel(ctx context.Context, model *models.Tenant) error {
	return svc.tenantDao.Update(ctx, model)
}

// Delete
func (svc *TenantService) Delete(ctx context.Context, id int64) error {
	return svc.tenantDao.Transaction(func() error {
		if err := svc.tenantDao.Delete(ctx, id); err != nil {
			return err
		}

		return svc.permissionRuleSvc.DeleteByTenantID(ctx, id)
	})
}
