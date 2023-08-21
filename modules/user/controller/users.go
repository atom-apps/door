package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type UserController struct {
	userSvc           *service.UserService
	permissionRuleSvc *service.PermissionRuleService
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
//	@Router			/v1/users/{id} [get]
func (c *UserController) Show(ctx *fiber.Ctx, id int64) (*dto.UserItem, error) {
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
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.UserItem}
//	@Router			/v1/users [get]
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

// Role users
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.UserItem}
//	@Router			/v1/users/roles/{role_id} [get]
func (c *UserController) Role(
	ctx *fiber.Ctx,
	roleID int64,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	var err error
	queryFilter.IDs, err = c.permissionRuleSvc.GetUserIDsOfRole(ctx.Context(), roleID)
	if err != nil {
		return nil, err
	}

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

// Tenant users
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.UserItem}
//	@Router			/v1/users/tenants/{tenant_id} [get]
func (c *UserController) Tenant(
	ctx *fiber.Ctx,
	tenantID int64,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	var err error
	queryFilter.IDs, err = c.permissionRuleSvc.GetUserIDsOfTenant(ctx.Context(), tenantID)
	if err != nil {
		return nil, err
	}

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
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Router			/v1/users [post]
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
//	@Router			/v1/users/{id} [put]
func (c *UserController) Update(ctx *fiber.Ctx, id int64, body *dto.UserForm) error {
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
//	@Router			/v1/users/{id} [delete]
func (c *UserController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.userSvc.Delete(ctx.Context(), id)
}
