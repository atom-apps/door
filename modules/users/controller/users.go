package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/modules/users/service"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type UserController struct {
	userSvc       *service.UserService
	permissionSvc *service.PermissionService
}

// Profile get current user info
//
//	@Summary		get current user info
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/v1/users/users/profile [get]
func (c *UserController) Profile(ctx *fiber.Ctx, claim *jwt.Claims) (*dto.UserItem, error) {
	item, err := c.userSvc.GetByID(ctx.Context(), claim.UserID)
	if err != nil {
		return nil, err
	}

	userItem := c.userSvc.DecorateItem(item, 0)
	userItem.Claims = &claim.BaseClaims
	return userItem, nil
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/v1/users/users/{id} [get]
func (c *UserController) Show(ctx *fiber.Ctx, id uint64) (*dto.UserItem, error) {
	item, err := c.userSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.userSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
//	@Param			pageFilter	query		ds.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		ds.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	ds.PageDataResponse{list=dto.UserItem}
//	@Router			/v1/users/users [get]
func (c *UserController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*ds.PageDataResponse, error) {
	items, total, err := c.userSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &ds.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.userSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Router			/v1/users/users [post]
func (c *UserController) Create(ctx *fiber.Ctx, body *dto.UserForm) error {
	return c.userSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"UserID"
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Failure		500		{string}	UserID
//	@Router			/v1/users/users/{id} [put]
func (c *UserController) Update(ctx *fiber.Ctx, id uint64, body *dto.UserForm) error {
	return c.userSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{string}	UserID
//	@Failure		500	{string}	UserID
//	@Router			/v1/users/users/{id} [delete]
func (c *UserController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.userSvc.Delete(ctx.Context(), id)
}

// ResetPassword
//
//	@Summary		ResetPassword
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"UserID"
//	@Success		200		{string}	UserID
//	@Router			/v1/users/users/{id}/reset-password [put]
func (c *UserController) ResetPassword(ctx *fiber.Ctx, id uint64) (string, error) {
	newPassword := common.RandomString(16)
	return newPassword, c.userSvc.ResetPassword(ctx.Context(), id, newPassword)
}
