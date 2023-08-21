package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/errorx"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/user/dao"
	"github.com/atom-apps/door/modules/user/dto"

	"github.com/jinzhu/copier"
)

// @provider
type RoleService struct {
	roleDao           *dao.RoleDao
	permissionRuleSvc *PermissionRuleService
}

func (svc *RoleService) DecorateItem(model *models.Role, id int) *dto.RoleItem {
	dtoItem := &dto.RoleItem{
		ID:          model.ID,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		ParentID:    model.ParentID,
		Parent:      nil,
	}

	if model.ParentID != 0 {
		pModel, err := svc.GetByID(context.Background(), model.ParentID)
		if err == nil {
			dtoItem.Parent = svc.DecorateItem(pModel, id)
		}
	}

	return dtoItem
}

func (svc *RoleService) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	return svc.roleDao.GetByID(ctx, id)
}

func (svc *RoleService) GetBySlug(ctx context.Context, slug string) (*models.Role, error) {
	return svc.roleDao.GetBySlug(ctx, slug)
}

func (svc *RoleService) GetByUserID(ctx context.Context, tenantID, userID int64) (*models.Role, error) {
	return svc.permissionRuleSvc.GetRoleOfTenantUser(ctx, tenantID, userID)
}

func (svc *RoleService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Role, error) {
	return svc.roleDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *RoleService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RoleListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Role, int64, error) {
	return svc.roleDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *RoleService) CreateFromModel(ctx context.Context, model *models.Role) error {
	if _, err := svc.GetBySlug(ctx, model.Slug); err != nil {
		return err
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
func (svc *RoleService) Update(ctx context.Context, id int64, body *dto.RoleForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
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
func (svc *RoleService) Delete(ctx context.Context, id int64) error {
	return svc.roleDao.Transaction(func() error {
		if err := svc.roleDao.Delete(ctx, id); err != nil {
			return err
		}

		// delete permission rules roles
		return svc.permissionRuleSvc.DeleteByRoleID(ctx, id)
	})
}

// AttachUsers
func (svc *RoleService) AttachUsers(ctx context.Context, tenantID, roleID int64, userIDs []int64) error {
	return svc.permissionRuleSvc.AddRoleUsers(ctx, tenantID, roleID, userIDs)
}

// DetachUsers
func (svc *RoleService) DetachUsers(ctx context.Context, tenantID, roleID int64, userIDs []int64) error {
	return svc.permissionRuleSvc.DeleteRoleUsers(ctx, tenantID, roleID, userIDs)
}
