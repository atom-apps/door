package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	systemDao "github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-providers/log"
	"github.com/samber/lo"

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
				Path:     route.Path,
				ParentID: route.ParentID,
				Metadata: route.Metadata,
				Children: svc.genTree(routes, route.ID),
			})
		}
	}
	return genRoutes
}

func (svc *PermissionService) TenantRoleSave(ctx context.Context, tenantID, roleID uint64, routeIDs []uint64) error {
	// 首先获取需要配置的所有路由，清理出来
	routes, err := svc.routeDao.GetByIDs(ctx, routeIDs)
	if err != nil {
		return err
	}

	routes = lo.FilterMap(routes, func(route *models.Route, _ int) (*models.Route, bool) {
		if route.ParentID == 0 {
			return route, true
		}

		parentID := route.ParentID
		for parentID != 0 {
			parentRoute, err := svc.routeDao.GetByID(ctx, parentID)
			if err != nil {
				return nil, false
			}
			parentID = parentRoute.ParentID

			if lo.Contains(routeIDs, parentRoute.ID) {
				return nil, false
			}

			if parentID != 0 {
				continue
			}
			break
		}

		return route, true
	})

	permissions := []*models.Permission{}
	lo.ForEach(routes, func(route *models.Route, _ int) {
		permissions = append(permissions, &models.Permission{TenantID: tenantID, RoleID: roleID, RouteID: route.ID})
	})

	return svc.permissionDao.Transaction(func() error {
		// 删除所有存在的路由
		if err := svc.permissionDao.DeleteByTenantRole(ctx, tenantID, roleID); err != nil {
			return err
		}
		// 创建新的路由
		return svc.permissionDao.CreateBatch(ctx, permissions, 100)
	})
}

// GetPermissionIDsByRoleID
func (svc *PermissionService) GetPermissionIDsByRoleID(ctx context.Context, tenantID, roleID uint64) ([]uint64, error) {
	return svc.permissionDao.GetRouteIDsByTenantIDAndRoleID(ctx, tenantID, roleID)
}

// CasbinPolicies
func (svc *PermissionService) CasbinPolicies(ctx context.Context) ([][]string, error) {
	all, err := svc.permissionDao.FindALl(ctx)
	if err != nil {
		return nil, err
	}

	routeIDs := lo.Map(all, func(item *models.Permission, _ int) uint64 { return item.RouteID })
	routes, err := svc.routeDao.GetByIDs(ctx, routeIDs)
	if err != nil {
		return nil, err
	}

	return svc.genCasbinPolicies(ctx, all, routes)
}

func (svc *PermissionService) CasbinPoliciesOfTenantRole(ctx context.Context, tenantID, roleID uint64) ([][]string, error) {
	all, err := svc.permissionDao.FindByTenantRole(ctx, tenantID, roleID)
	if err != nil {
		return nil, err
	}

	routeIDs := lo.Map(all, func(item *models.Permission, _ int) uint64 { return item.RouteID })
	routes, err := svc.routeDao.GetByIDs(ctx, routeIDs)
	if err != nil {
		return nil, err
	}

	return svc.genCasbinPoliciesForPage(ctx, all, routes)
}

func (svc *PermissionService) genCasbinPoliciesForPage(ctx context.Context, ms []*models.Permission, routes []*models.Route) ([][]string, error) {
	routeMap := lo.KeyBy(routes, func(item *models.Route) uint64 {
		return item.ID
	})

	policies := [][]string{}
	lo.ForEach(ms, func(item *models.Permission, _ int) {
		if m, ok := routeMap[item.RouteID]; ok {
			policies = append(policies, []string{
				fmt.Sprintf("role:%d", item.RoleID),
				fmt.Sprintf("tenant:%d", item.TenantID),
				m.Path,
				"ANY",
			})
		}
	})
	return policies, nil
}

func (svc *PermissionService) genCasbinPolicies(ctx context.Context, ms []*models.Permission, routes []*models.Route) ([][]string, error) {
	routeMap := lo.KeyBy(routes, func(item *models.Route) uint64 {
		return item.ID
	})

	policies := [][]string{}
	lo.ForEach(ms, func(item *models.Permission, _ int) {
		if m, ok := routeMap[item.RouteID]; ok {
			lo.ForEach(m.API.Data, func(api string, _ int) {
				apis := strings.Split(api, "#")

				method, path := strings.ToUpper(apis[0]), apis[1]

				policies = append(policies, []string{
					fmt.Sprintf("role:%d", item.RoleID),
					fmt.Sprintf("tenant:%d", item.TenantID),
					path,
					method,
				})
			})
		}
	})
	return policies, nil
}
