package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/atom-apps/door/modules/systems/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type MenuController struct {
	menuSvc *service.MenuService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			Menu
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"MenuID"
//	@Success		200	{object}	dto.MenuItem
//	@Router			/v1/systems/menus/{id} [get]
func (c *MenuController) Show(ctx *fiber.Ctx, id uint64) (*dto.MenuItem, error) {
	item, err := c.menuSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.menuSvc.DecorateItem(item, 0), nil
}

// Tree
//
//	@Summary		Tree
//	@Tags			Menu
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"MenuID"
//	@Success		200	{object}	dto.MenuItem
//	@Router			/v1/systems/menus/{id}/tree [get]
func (c *MenuController) ShowTree(ctx *fiber.Ctx, id uint64) ([]*dto.MenuTreeItem, error) {
	return c.menuSvc.GetGroupTree(ctx.Context(), id)
}

// List list by query filter
//
//	@Summary		List
//	@Tags			Menu
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.MenuListQueryFilter	true	"MenuListQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.MenuItem}
//	@Router			/v1/systems/menus [get]
func (c *MenuController) List(ctx *fiber.Ctx, queryFilter *dto.MenuListQueryFilter, sortFilter *common.SortQueryFilter) ([]*dto.MenuItem, error) {
	items, err := c.menuSvc.FindGroupByQueryFilter(ctx.Context(), queryFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return lo.Map(items, c.menuSvc.DecorateItem), nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			Menu
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.MenuForm	true	"MenuForm"
//	@Success		200		{string}	MenuID
//	@Router			/v1/systems/menus [post]
func (c *MenuController) Create(ctx *fiber.Ctx, body *dto.MenuForm) error {
	body.ParentID = 0
	return c.menuSvc.Create(ctx.Context(), body, 0)
}

// CreateSub
//
//	@Summary		CreateSub
//	@Tags			Menu
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.MenuForm	true	"MenuForm"
//	@Success		200		{string}	MenuID
//	@Router			/v1/systems/menus/{menu_id}/sub [post]
func (c *MenuController) CreateSub(ctx *fiber.Ctx, menuID uint64, body *dto.MenuForm) error {
	return c.menuSvc.Create(ctx.Context(), body, menuID)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			Menu
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"MenuID"
//	@Param			body	body		dto.MenuForm	true	"MenuForm"
//	@Success		200		{string}	MenuID
//	@Router			/v1/systems/menus/{id} [put]
func (c *MenuController) Update(ctx *fiber.Ctx, id uint64, body *dto.MenuForm) error {
	return c.menuSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			Menu
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"MenuID"
//	@Success		200	{string}	MenuID
//	@Router			/v1/systems/menus/{id} [delete]
func (c *MenuController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.menuSvc.Delete(ctx.Context(), id)
}
