package service

import (
	"context"
	"errors"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/common/errorx"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-providers/log"
	"github.com/samber/lo"

	"github.com/jinzhu/copier"
)

// @provider
type RoleService struct {
	roleDao           *dao.RoleDao
	userTenantRoleSvc *UserTenantRoleService
	permissionSvc     *PermissionService
}

func (svc *RoleService) DecorateItem(model *models.Role, id int) *dto.RoleItem {
	userAmount, err := svc.userTenantRoleSvc.GetUserAmountOfRole(context.Background(), model.ID)
	if err != nil {
		log.Warnf("get user amount of role %d failed: %v", model.ID, err)
	}

	tenants, err := svc.userTenantRoleSvc.GetTenantsByRoleID(context.Background(), model.ID)
	if err != nil {
		log.Warnf("get tenants of role %d failed: %v", model.ID, err)
	}

	permissions := make(map[uint64][]uint64)
	lo.ForEach(tenants, func(item *models.Tenant, _ int) {
		routes, err := svc.permissionSvc.permissionDao.GetRouteIDsByTenantIDAndRoleID(context.Background(), item.ID, model.ID)
		if err != nil {
			return
		}
		permissions[item.ID] = routes
	})

	dtoItem := &dto.RoleItem{
		ID:          model.ID,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		ParentID:    model.ParentID,
		Parent:      nil,
		UserAmount:  userAmount,
		Tenants:     tenants,
		Permissions: permissions,
	}

	if model.ParentID != 0 {
		pModel, err := svc.GetByID(context.Background(), model.ParentID)
		if err == nil {
			dtoItem.Parent = svc.DecorateItem(pModel, id)
		}
	}

	return dtoItem
}

func (svc *RoleService) GetByID(ctx context.Context, id uint64) (*models.Role, error) {
	return svc.roleDao.GetByID(ctx, id)
}

func (svc *RoleService) GetBySlug(ctx context.Context, slug string) (*models.Role, error) {
	return svc.roleDao.GetBySlug(ctx, slug)
}

func (svc *RoleService) GetByUserID(ctx context.Context, tenantID, userID uint64) (*models.Role, error) {
	return svc.userTenantRoleSvc.GetRoleOfTenantUser(ctx, tenantID, userID)
}

func (svc *RoleService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Role, error) {
	return svc.roleDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *RoleService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.Role, int64, error) {
	return svc.roleDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *RoleService) CreateFromModel(ctx context.Context, model *models.Role) error {
	if _, err := svc.GetBySlug(ctx, model.Slug); err == nil {
		return errors.New("slug exists")
	}
	return svc.roleDao.Create(ctx, model)
}

// Create
func (svc *RoleService) Create(ctx context.Context, body *dto.RoleForm) error {
	model := &models.Role{}
	_ = copier.Copy(model, body)
	return svc.CreateFromModel(ctx, model)
}

// Update
func (svc *RoleService) Update(ctx context.Context, id uint64, body *dto.RoleForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if body.ParentID == model.ID {
		body.ParentID = 0
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.UpdateFromModel(ctx, model)
}

// UpdateFromModel
func (svc *RoleService) UpdateFromModel(ctx context.Context, model *models.Role) error {
	if svc.roleDao.SlugExists(ctx, model) {
		return errorx.ErrRecordExists
	}
	return svc.roleDao.Update(ctx, model)
}

// Delete
func (svc *RoleService) Delete(ctx context.Context, id uint64) error {
	return svc.roleDao.Transaction(func() error {
		if err := svc.roleDao.Delete(ctx, id); err != nil {
			return err
		}

		// delete permission rules roles
		if err := svc.userTenantRoleSvc.DeleteByRoleID(ctx, id); err != nil {
			return err
		}

		if err := svc.permissionSvc.DeleteByRoleID(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
