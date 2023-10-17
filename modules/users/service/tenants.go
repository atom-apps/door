package service

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-providers/log"

	"github.com/jinzhu/copier"
)

// @provider
type TenantService struct {
	tenantDao         *dao.TenantDao
	permissionSvc     *PermissionService
	userTenantRoleSvc *UserTenantRoleService
}

func (svc *TenantService) DecorateItem(model *models.Tenant, id int) *dto.TenantItem {
	userAmount, err := svc.userTenantRoleSvc.GetUserAmountOfTenant(context.Background(), model.ID)
	if err != nil {
		log.Warnf("get user amount of tenant %d failed: %v", model.ID, err)
	}

	roles, err := svc.userTenantRoleSvc.GetRolesByTenantID(context.Background(), model.ID)
	if err != nil {
		log.Warnf("get roles of tenant %d failed: %v", model.ID, err)
	}

	return &dto.TenantItem{
		ID:          model.ID,
		CreatedAt:   model.CreatedAt,
		Name:        model.Name,
		Description: model.Description,
		Meta:        model.Meta,
		UserAmount:  userAmount,
		Roles:       roles,
	}
}

func (svc *TenantService) GetByID(ctx context.Context, id uint64) (*models.Tenant, error) {
	return svc.tenantDao.GetByID(ctx, id)
}

func (svc *TenantService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Tenant, error) {
	return svc.tenantDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *TenantService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TenantListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
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
func (svc *TenantService) Update(ctx context.Context, id uint64, body *dto.TenantForm) error {
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
func (svc *TenantService) Delete(ctx context.Context, id uint64) error {
	return svc.tenantDao.Transaction(func() error {
		if err := svc.tenantDao.Delete(ctx, id); err != nil {
			return err
		}

		if err := svc.userTenantRoleSvc.DeleteByTenantID(ctx, id); err != nil {
			return err
		}

		if err := svc.permissionSvc.DeleteByTenantID(ctx, id); err != nil {
			return err
		}

		return nil
	})
}
