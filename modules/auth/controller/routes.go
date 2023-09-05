package controller

import (
	"encoding/json"

	"github.com/atom-apps/door/docs"
	"github.com/atom-apps/door/modules/auth/dto"
	systemDto "github.com/atom-apps/door/modules/systems/dto"
	systemSvc "github.com/atom-apps/door/modules/systems/service"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
)

// @provider
type RoutesController struct {
	routeSvc *systemSvc.RouteService
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
//	@Success	200	{array}	dto.RouteItem
//	@Router		/v1/auth/pages [get]
func (c *RoutesController) Pages(ctx *fiber.Ctx, claim *jwt.Claims) ([]*systemDto.RouteItem, error) {
	return c.routeSvc.Tree(ctx.Context(), 0)
}
