package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/atom-apps/door/modules/systems/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type LocationController struct {
	locationSvc *service.LocationService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"LocationID"
//	@Success		200	{object}	dto.LocationItem
//	@Router			/v1/systems/locations/{id} [get]
func (c *LocationController) Show(ctx *fiber.Ctx, id uint64) (*dto.LocationItem, error) {
	item, err := c.locationSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.locationSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.LocationListQueryFilter	true	"LocationListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.LocationItem}
//	@Router			/v1/systems/locations [get]
func (c *LocationController) List(
	ctx *fiber.Ctx,
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
//	@Summary		Create
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.LocationForm	true	"LocationForm"
//	@Success		200		{string}	LocationID
//	@Router			/v1/systems/locations [post]
func (c *LocationController) Create(ctx *fiber.Ctx, body *dto.LocationForm) error {
	return c.locationSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"LocationID"
//	@Param			body	body		dto.LocationForm	true	"LocationForm"
//	@Success		200		{string}	LocationID
//	@Router			/v1/systems/locations/{id} [put]
func (c *LocationController) Update(ctx *fiber.Ctx, id uint64, body *dto.LocationForm) error {
	return c.locationSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"LocationID"
//	@Success		200	{string}	LocationID
//	@Router			/v1/systems/locations/{id} [delete]
func (c *LocationController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.locationSvc.Delete(ctx.Context(), id)
}

// Provinces
//
//	@Summary		省
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}	dto.LocationItem
//	@Router			/v1/systems/locations/provinces [get]
func (c *LocationController) Provinces(ctx *fiber.Ctx) ([]*dto.LocationItem, error) {
	items, err := c.locationSvc.Provinces(ctx.Context())
	if err != nil {
		return nil, err
	}

	return lo.Map(items, c.locationSvc.DecorateItem), nil
}

// Cities
//
//	@Summary		市
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}	dto.LocationItem
//	@Router			/v1/systems/locations/provinces/{province}/cities [get]
func (c *LocationController) Cities(ctx *fiber.Ctx, province string) ([]*dto.LocationItem, error) {
	items, err := c.locationSvc.Cities(ctx.Context(), province)
	if err != nil {
		return nil, err
	}

	return lo.Map(items, c.locationSvc.DecorateItem), nil
}

// Areas
//
//	@Summary		地区
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}	dto.LocationItem
//	@Router			/v1/systems/locations/provinces/{province}/cities/{city}/areas [get]
func (c *LocationController) Areas(ctx *fiber.Ctx, province, city string) ([]*dto.LocationItem, error) {
	items, err := c.locationSvc.Areas(ctx.Context(), province, city)
	if err != nil {
		return nil, err
	}

	return lo.Map(items, c.locationSvc.DecorateItem), nil
}

// Towns
//
//	@Summary		街道
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}	dto.LocationItem
//	@Router			/v1/systems/locations/provinces/{province}/cities/{city}/areas/{area}/town [get]
func (c *LocationController) Towns(ctx *fiber.Ctx, province, city, area string) ([]*dto.LocationItem, error) {
	items, err := c.locationSvc.Towns(ctx.Context(), province, city, area)
	if err != nil {
		return nil, err
	}

	return lo.Map(items, c.locationSvc.DecorateItem), nil
}

// Location
//
//	@Summary		地区全称
//	@Tags			区域
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	dto.LocationDetail
//	@Router			/v1/systems/locations/{code}-{town} [get]
func (c *LocationController) Location(ctx *fiber.Ctx, code, town string) (*dto.LocationDetail, error) {
	return c.locationSvc.GetByCodeTown(ctx.Context(), code, town)
}
