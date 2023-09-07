package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/atom-apps/door/modules/systems/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type DictionaryController struct {
	dictionarySvc *service.DictionaryService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DictionaryID"
//	@Success		200	{object}	dto.DictionaryItem
//	@Router			/v1/systems/dictionaries/{id} [get]
func (c *DictionaryController) Show(ctx *fiber.Ctx, id uint64) (*dto.DictionaryItem, error) {
	item, err := c.dictionarySvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.dictionarySvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.DictionaryListQueryFilter	true	"DictionaryListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.DictionaryItem}
//	@Router			/v1/systems/dictionaries [get]
func (c *DictionaryController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.DictionaryListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.dictionarySvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.dictionarySvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.DictionaryForm	true	"DictionaryForm"
//	@Success		200		{string}	DictionaryID
//	@Router			/v1/systems/dictionaries [post]
func (c *DictionaryController) Create(ctx *fiber.Ctx, body *dto.DictionaryForm) error {
	return c.dictionarySvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"DictionaryID"
//	@Param			body	body		dto.DictionaryForm	true	"DictionaryForm"
//	@Success		200		{string}	DictionaryID
//	@Router			/v1/systems/dictionaries/{id} [put]
func (c *DictionaryController) Update(ctx *fiber.Ctx, id uint64, body *dto.DictionaryForm) error {
	return c.dictionarySvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DictionaryID"
//	@Success		200	{string}	DictionaryID
//	@Router			/v1/systems/dictionaries/{id} [delete]
func (c *DictionaryController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.dictionarySvc.Delete(ctx.Context(), id)
}
