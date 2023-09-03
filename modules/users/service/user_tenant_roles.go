package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/samber/lo"

	"github.com/jinzhu/copier"
)

// @provider
type UserTenantRoleService struct {
	userTenantRoleDao *dao.UserTenantRoleDao
	roleDao           *dao.RoleDao
	tenantDao         *dao.TenantDao
}

func (svc *UserTenantRoleService) DecorateItem(model *models.UserTenantRole, id int) *dto.UserTenantRoleItem {
	var dtoItem *dto.UserTenantRoleItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *UserTenantRoleService) GetByID(ctx context.Context, id uint64) (*models.UserTenantRole, error) {
	return svc.userTenantRoleDao.GetByID(ctx, id)
}

func (svc *UserTenantRoleService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserTenantRoleListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserTenantRole, error) {
	return svc.userTenantRoleDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *UserTenantRoleService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserTenantRoleListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.UserTenantRole, int64, error) {
	return svc.userTenantRoleDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *UserTenantRoleService) CreateFromModel(ctx context.Context, model *models.UserTenantRole) error {
	return svc.userTenantRoleDao.Create(ctx, model)
}

// Create
func (svc *UserTenantRoleService) Create(ctx context.Context, body *dto.UserTenantRoleForm) error {
	model := &models.UserTenantRole{}
	_ = copier.Copy(model, body)
	return svc.userTenantRoleDao.Create(ctx, model)
}

// Update
func (svc *UserTenantRoleService) Update(ctx context.Context, id uint64, body *dto.UserTenantRoleForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.userTenantRoleDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *UserTenantRoleService) UpdateFromModel(ctx context.Context, model *models.UserTenantRole) error {
	return svc.userTenantRoleDao.Update(ctx, model)
}

// Delete
func (svc *UserTenantRoleService) Delete(ctx context.Context, id uint64) error {
	return svc.userTenantRoleDao.Delete(ctx, id)
}

func (svc *UserTenantRoleService) AttachUsers(ctx context.Context, tenantID, roleID uint64, users []uint64) error {
	ms := lo.FilterMap(users, func(userID uint64, _ int) (*models.UserTenantRole, bool) {
		if ok, _ := svc.userTenantRoleDao.TenantHasRole(ctx, tenantID, userID); ok {
			return nil, false
		}
		return &models.UserTenantRole{UserID: userID, TenantID: tenantID, RoleID: roleID}, true
	})

	if len(ms) == 0 {
		return nil
	}

	return svc.userTenantRoleDao.CreateInBatch(ctx, ms)
}

func (svc *UserTenantRoleService) DetachUsers(ctx context.Context, tenantID, roleID uint64, users []uint64) error {
	ms := lo.Map(users, func(item uint64, _ int) *models.UserTenantRole {
		return &models.UserTenantRole{UserID: item, TenantID: tenantID, RoleID: roleID}
	})

	return svc.userTenantRoleDao.DeleteInBatch(ctx, ms)
}

func (svc *UserTenantRoleService) DeleteTenant(ctx context.Context, tenantID uint64) error {
	return svc.userTenantRoleDao.DeleteByTenantID(ctx, tenantID)
}

func (svc *UserTenantRoleService) DeleteRole(ctx context.Context, roleID uint64) error {
	return svc.userTenantRoleDao.DeleteByRoleID(ctx, roleID)
}

// GetUserAmountOfRole
func (svc *UserTenantRoleService) GetUserAmountOfRole(ctx context.Context, roleID uint64) (int64, error) {
	return svc.userTenantRoleDao.GetUserAmountOfRole(ctx, roleID)
}

// GetUserAmountOfTenant
func (svc *UserTenantRoleService) GetUserAmountOfTenant(ctx context.Context, tenantID uint64) (int64, error) {
	return svc.userTenantRoleDao.GetUserAmountOfTenant(ctx, tenantID)
}

// GetRoleOfTenantUser
func (svc *UserTenantRoleService) GetRoleOfTenantUser(ctx context.Context, tenantID, userID uint64) (*models.Role, error) {
	m, err := svc.userTenantRoleDao.GetRoleOfTenantUser(ctx, tenantID, userID)
	if err != nil {
		return nil, err
	}

	return svc.roleDao.GetByID(ctx, m.RoleID)
}

// DeleteByRoleID
func (svc *UserTenantRoleService) DeleteByRoleID(ctx context.Context, roleID uint64) error {
	return svc.userTenantRoleDao.DeleteByRoleID(ctx, roleID)
}

// DeleteByTenantID
func (svc *UserTenantRoleService) DeleteByTenantID(ctx context.Context, tenantID uint64) error {
	return svc.userTenantRoleDao.DeleteByTenantID(ctx, tenantID)
}

// GetTenantsByUserID
func (svc *UserTenantRoleService) GetTenantsByUserID(ctx context.Context, userID uint64) ([]*models.Tenant, error) {
	ms, err := svc.userTenantRoleDao.GetTenantsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	ids := lo.Map(ms, func(item *models.UserTenantRole, _ int) uint64 {
		return item.TenantID
	})

	return svc.tenantDao.FindByIDs(ctx, ids)
}

// DeleteByUserID
func (svc *UserTenantRoleService) DeleteByUserID(ctx context.Context, userID uint64) error {
	return svc.userTenantRoleDao.DeleteByUserID(ctx, userID)
}

// TenantHasRole
func (svc *UserTenantRoleService) TenantHasRole(ctx context.Context, tenantID, userID uint64) (bool, error) {
	return svc.userTenantRoleDao.TenantHasRole(ctx, tenantID, userID)
}
