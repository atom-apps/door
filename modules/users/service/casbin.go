package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/atom-apps/door/database/models"
	systemDao "github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/users/dao"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/log"
	"github.com/samber/lo"
)

// @provider
type CasbinService struct {
	casbin            *casbin.Casbin
	permissionDao     *dao.PermissionDao
	userTenantRoleDao *dao.UserTenantRoleDao
	tenantDao         *dao.TenantDao
	roleDao           *dao.RoleDao
	routeDao          *systemDao.RouteDao
}

// CasbinPolicies
func (svc *CasbinService) CasbinPolicies(ctx context.Context) ([][]string, error) {
	all, err := svc.permissionDao.FindALl(ctx)
	if err != nil {
		return nil, err
	}

	routeIDs := lo.Map(all, func(item *models.Permission, _ int) uint64 { return item.RouteID })
	routes, err := svc.routeDao.GetByIDsWithParents(ctx, routeIDs)
	if err != nil {
		return nil, err
	}

	return svc.genCasbinPolicies(ctx, all, routes)
}

func (svc *CasbinService) genCasbinPolicies(ctx context.Context, ms []*models.Permission, routes []*models.Route) ([][]string, error) {
	routeMap := lo.KeyBy(routes, func(item *models.Route) uint64 {
		return item.ID
	})

	policies := [][]string{}
	lo.ForEach(ms, func(item *models.Permission, _ int) {
		if m, ok := routeMap[item.RouteID]; ok {
			lo.ForEach(m.API, func(api string, _ int) {
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

// GroupPolicy
func (svc *CasbinService) CasbinGroups(ctx context.Context) ([][]string, error) {
	all, err := svc.userTenantRoleDao.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return svc.genCasbinGroups(ctx, all)
}

// GroupPolicy
func (svc *CasbinService) CasbinUserGroups(ctx context.Context, userID uint64) ([][]string, error) {
	all, err := svc.userTenantRoleDao.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return svc.genCasbinGroups(ctx, all)
}

func (svc *CasbinService) genCasbinGroups(ctx context.Context, ms []*models.UserTenantRole) ([][]string, error) {
	policies := [][]string{}

	// g, user, role, tenant
	for _, item := range ms {
		policies = append(policies, []string{
			fmt.Sprintf("%d", item.UserID),
			fmt.Sprintf("role:%d", item.RoleID),
			fmt.Sprintf("tenant:%d", item.TenantID),
		})
	}

	return policies, nil
}

func (svc *CasbinService) Reload() error {
	if err := svc.casbin.Clear(); err != nil {
		return err
	}

	groups, err := svc.CasbinGroups(context.Background())
	if err != nil {
		return err
	}

	if _, err := svc.casbin.LoadGroups(groups); err != nil {
		return err
	}
	log.Infof("load groups: %d", len(groups))

	// permissions
	policies, err := svc.CasbinPolicies(context.Background())
	if err != nil {
		return err
	}
	if _, err := svc.casbin.LoadPolicies(policies); err != nil {
		return err
	}
	log.Infof("load policies: %d", len(policies))

	return nil
}
