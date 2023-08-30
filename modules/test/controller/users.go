package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/test/dto"
	"github.com/atom-apps/door/modules/test/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type UserController struct {
	userSvc *service.UserService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/v1/users/roles/{id} [get]
func (c *UserController) Show(ctx *fiber.Ctx, id int64) (*dto.UserItem, error) {
	item, err := c.userSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.userSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.UserItem}
//	@Router			/v1/users/roles [get]
func (c *UserController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.userSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.userSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Router			/v1/users/roles [post]
func (c *UserController) Create(ctx *fiber.Ctx, body *dto.UserForm) error {
	return c.userSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"UserID"
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Router			/v1/users/roles/{id} [put]
func (c *UserController) Update(ctx *fiber.Ctx, id int64, body *dto.UserForm) error {
	return c.userSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{string}	UserID
//	@Router			/v1/users/roles/{id} [delete]
func (c *UserController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.userSvc.Delete(ctx.Context(), id)
}
