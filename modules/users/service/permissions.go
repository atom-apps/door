package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	systemDao "github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-providers/log"

	"github.com/jinzhu/copier"
)

// @provider
type PermissionService struct {
	permissionDao *dao.PermissionDao
	tenantDao     *dao.TenantDao
	roleDao       *dao.RoleDao
	routeDao      *systemDao.RouteDao
}

func (svc *PermissionService) DecorateItem(model *models.Permission, id int) *dto.PermissionItem {
	role, err := svc.roleDao.GetByID(context.Background(), model.RoleID)
	if err != nil {
		log.Warn("get role failed by id: %d", model.RoleID)
	}

	tenant, err := svc.tenantDao.GetByID(context.Background(), model.TenantID)
	if err != nil {
		log.Warn("get tenant failed by id: %d", model.TenantID)
	}

	return &dto.PermissionItem{
		ID:       model.ID,
		TenantID: model.TenantID,
		Tenant:   tenant,
		RoleID:   model.RoleID,
		Role:     role,
		Path:     model.Path,
		Action:   model.Action,
	}
}

func (svc *PermissionService) GetByID(ctx context.Context, id uint64) (*models.Permission, error) {
	return svc.permissionDao.GetByID(ctx, id)
}

func (svc *PermissionService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.PermissionListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Permission, error) {
	return svc.permissionDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *PermissionService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.PermissionListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Permission, int64, error) {
	return svc.permissionDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *PermissionService) CreateFromModel(ctx context.Context, model *models.Permission) error {
	return svc.permissionDao.Create(ctx, model)
}

// Create
func (svc *PermissionService) Create(ctx context.Context, body *dto.PermissionForm) error {
	model := &models.Permission{}
	_ = copier.Copy(model, body)
	return svc.permissionDao.Create(ctx, model)
}

// Update
func (svc *PermissionService) Update(ctx context.Context, id uint64, body *dto.PermissionForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.permissionDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *PermissionService) UpdateFromModel(ctx context.Context, model *models.Permission) error {
	return svc.permissionDao.Update(ctx, model)
}

// Delete
func (svc *PermissionService) Delete(ctx context.Context, id uint64) error {
	return svc.permissionDao.Delete(ctx, id)
}

func (svc *PermissionService) DeleteByTenantID(ctx context.Context, tenantID uint64) error {
	return svc.permissionDao.DeleteByTenantID(ctx, tenantID)
}

func (svc *PermissionService) DeleteByRoleID(ctx context.Context, roleID uint64) error {
	return svc.permissionDao.DeleteByRoleID(ctx, roleID)
}

func (svc *PermissionService) Tree(ctx context.Context) ([]*dto.PermissionTree, error) {
	routes, err := svc.routeDao.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return svc.genTree(routes, 0), nil
}

func (svc *PermissionService) genTree(routes []*models.Route, parentID uint64) []*dto.PermissionTree {
	genRoutes := []*dto.PermissionTree{}
	for _, route := range routes {
		if route.ParentID == parentID {
			genRoutes = append(genRoutes, &dto.PermissionTree{
				ID:       route.ID,
				Name:     route.Name,
				Method:   route.Method,
				Path:     route.Path,
				ParentID: route.ParentID,
				Metadata: route.Metadata,
				Children: svc.genTree(routes, route.ID),
			})
		}
	}
	return genRoutes
}
