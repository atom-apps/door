package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/tools/dto"
	"github.com/atom-apps/door/modules/tools/service"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type LocationController struct {
	locationSvc *service.LocationService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			Tool
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"LocationID"
//	@Success		200	{object}	dto.LocationItem
//	@Router			/v1/tools/locations/{id} [get]
func (c *LocationController) Show(ctx *fiber.Ctx, claim *jwt.Claims, id int64) (*dto.LocationItem, error) {
	item, err := c.locationSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.locationSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			Tool
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.LocationListQueryFilter	true	"LocationListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter		true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter		true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.LocationItem}
//	@Router			/v1/tools/locations [get]
func (c *LocationController) List(
	ctx *fiber.Ctx,
	claim *jwt.Claims,
	queryFilter *dto.LocationListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.locationSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.locationSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			Tool
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.LocationForm	true	"LocationForm"
//	@Success		200		{string}	LocationID
//	@Router			/v1/tools/locations [post]
func (c *LocationController) Create(ctx *fiber.Ctx, claim *jwt.Claims, body *dto.LocationForm) error {
	return c.locationSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			Tool
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"LocationID"
//	@Param			body	body		dto.LocationForm	true	"LocationForm"
//	@Success		200		{string}	LocationID
//	@Failure		500		{string}	LocationID
//	@Router			/v1/tools/locations/{id} [put]
func (c *LocationController) Update(ctx *fiber.Ctx, claim *jwt.Claims, id int64, body *dto.LocationForm) error {
	return c.locationSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			Tool
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"LocationID"
//	@Success		200	{string}	LocationID
//	@Failure		500	{string}	LocationID
//	@Router			/v1/tools/locations/{id} [delete]
func (c *LocationController) Delete(ctx *fiber.Ctx, claim *jwt.Claims, id int64) error {
	return c.locationSvc.Delete(ctx.Context(), id)
}
