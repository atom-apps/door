package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/test/dto"
	"github.com/atom-apps/door/modules/test/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type RoleController struct {
	roleSvc *service.RoleService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"RoleID"
//	@Success		200	{object}	dto.RoleItem
//	@Router			/v1/users/roles/{id} [get]
func (c *RoleController) Show(ctx *fiber.Ctx, id int64) (*dto.RoleItem, error) {
	item, err := c.roleSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.roleSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.RoleListQueryFilter	true	"RoleListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.RoleItem}
//	@Router			/v1/users/roles [get]
func (c *RoleController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.RoleListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.roleSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.roleSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.RoleForm	true	"RoleForm"
//	@Success		200		{string}	RoleID
//	@Router			/v1/users/roles [post]
func (c *RoleController) Create(ctx *fiber.Ctx, body *dto.RoleForm) error {
	return c.roleSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"RoleID"
//	@Param			body	body		dto.RoleForm	true	"RoleForm"
//	@Success		200		{string}	RoleID
//	@Router			/v1/users/roles/{id} [put]
func (c *RoleController) Update(ctx *fiber.Ctx, id int64, body *dto.RoleForm) error {
	return c.roleSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"RoleID"
//	@Success		200	{string}	RoleID
//	@Router			/v1/users/roles/{id} [delete]
func (c *RoleController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.roleSvc.Delete(ctx.Context(), id)
}
