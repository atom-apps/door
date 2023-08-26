package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/common/errorx"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/atom-apps/door/modules/systems/service"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type RouteController struct {
	routeSvc *service.RouteService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			Systems
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"RouteID"
//	@Success		200	{object}	dto.RouteItem
//	@Router			/v1/routes/{id} [get]
func (c *RouteController) Show(ctx *fiber.Ctx, claim *jwt.Claims, id int64) (*dto.RouteItem, error) {
	if claim.IsAdmin() {
		return nil, errorx.ErrForbidden
	}
	item, err := c.routeSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.routeSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			Systems
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.RouteListQueryFilter	true	"RouteListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter		true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter		true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.RouteItem}
//	@Router			/v1/routes [get]
func (c *RouteController) List(
	ctx *fiber.Ctx,
	claim *jwt.Claims,
	queryFilter *dto.RouteListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	if claim.IsAdmin() {
		return nil, errorx.ErrForbidden
	}
	items, total, err := c.routeSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.routeSvc.DecorateItem),
	}, nil
}

// Pages get page routes
//
//	@Summary	获取页面路由
//	@Tags		Systems
//	@Accept		json
//	@Produce	json
//	@Param		queryFilter	query	dto.RouteListQueryFilter	true	"RouteListQueryFilter"
//	@Success	200			{array}	dto.RouteItem
//	@Router		/v1/routes/type/{route_type} [get]
func (c *RouteController) Pages(ctx *fiber.Ctx, claim *jwt.Claims, routeType string) ([]*dto.RouteItem, error) {
	typ, err := consts.ParseRouteType(routeType)
	if err != nil {
		return nil, err
	}

	return c.routeSvc.Tree(ctx.Context(), typ, 0)
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			Systems
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.RouteForm	true	"RouteForm"
//	@Success		200		{string}	RouteID
//	@Router			/v1/routes [post]
func (c *RouteController) Create(ctx *fiber.Ctx, claim *jwt.Claims, body *dto.RouteForm) error {
	if claim.IsAdmin() {
		return errorx.ErrForbidden
	}
	return c.routeSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			Systems
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"RouteID"
//	@Param			body	body		dto.RouteForm	true	"RouteForm"
//	@Success		200		{string}	RouteID
//	@Failure		500		{string}	RouteID
//	@Router			/v1/routes/{id} [put]
func (c *RouteController) Update(ctx *fiber.Ctx, claim *jwt.Claims, id int64, body *dto.RouteForm) error {
	if claim.IsAdmin() {
		return errorx.ErrForbidden
	}
	return c.routeSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			Systems
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"RouteID"
//	@Success		200	{string}	RouteID
//	@Failure		500	{string}	RouteID
//	@Router			/v1/routes/{id} [delete]
func (c *RouteController) Delete(ctx *fiber.Ctx, claim *jwt.Claims, id int64) error {
	if claim.IsAdmin() {
		return errorx.ErrForbidden
	}
	return c.routeSvc.Delete(ctx.Context(), id)
}
