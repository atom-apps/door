package controller

import (
	"encoding/json"

	"github.com/atom-apps/door/common/model"
	"github.com/atom-apps/door/docs"
	"github.com/atom-apps/door/modules/auth/dto"
	systemSvc "github.com/atom-apps/door/modules/systems/service"
	userSvc "github.com/atom-apps/door/modules/users/service"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
)

// @provider
type RoutesController struct {
	routeSvc      *systemSvc.RouteService
	permissionSvc *userSvc.PermissionService
	casbinSvc     *userSvc.CasbinService
}

// List
//
//	@Summary		Signup
//	@Description	Signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.ExchangeTokenByCodeForm
//	@Router			/v1/auth/routes [get]
func (c *RoutesController) List(ctx *fiber.Ctx) ([]*dto.Route, error) {
	var doc *dto.SwaggerDoc
	err := json.Unmarshal([]byte(docs.SwaggerSpec), &doc)
	if err != nil {
		return nil, err
	}
	return doc.ToRoues(), nil
}

// Pages get page routes
//
//	@Summary	获取页面路由
//	@Tags		Systems
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}	model.RouteItem
//	@Router		/v1/auth/pages [get]
func (c *RoutesController) Pages(ctx *fiber.Ctx, claim *jwt.Claims) ([]*model.RouteItem, error) {
	if claim.IsSuperAdmin() {
		return c.permissionSvc.Pages(ctx.Context())
	}
	return c.permissionSvc.PagesOfTenantRole(ctx.Context(), claim.TenantID, claim.RoleID)
}

// Test
//
//	@Summary	获取页面路由
//	@Tags		Systems
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}	dto.RouteItem
//	@Router		/v1/auth/test [get]
func (c *RoutesController) Test(ctx *fiber.Ctx) ([][]string, error) {
	// return c.permissionSvc.CasbinPolicies(ctx.Context())
	return c.casbinSvc.CasbinGroups(ctx.Context())
}
