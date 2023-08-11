package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type TenantUserController struct {
	tenantUserSvc *service.TenantUserService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"TenantUserID"
//	@Success		200	{object}	dto.TenantUserItem
//	@Router			/tenant_users/{id} [get]
func (c *TenantUserController) Show(ctx *fiber.Ctx, id int64) (*dto.TenantUserItem, error) {
	item, err := c.tenantUserSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.tenantUserSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.TenantUserListQueryFilter	true	"TenantUserListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.TenantUserItem}
//	@Router			/tenant_users [get]
func (c *TenantUserController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.TenantUserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.tenantUserSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.tenantUserSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.TenantUserForm	true	"TenantUserForm"
//	@Success		200		{string}	TenantUserID
//	@Router			/tenant_users [post]
func (c *TenantUserController) Create(ctx *fiber.Ctx, body *dto.TenantUserForm) error {
	return c.tenantUserSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"TenantUserID"
//	@Param			body	body		dto.TenantUserForm	true	"TenantUserForm"
//	@Success		200		{string}	TenantUserID
//	@Failure		500		{string}	TenantUserID
//	@Router			/tenant_users/{id} [put]
func (c *TenantUserController) Update(ctx *fiber.Ctx, id int64, body *dto.TenantUserForm) error {
	return c.tenantUserSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"TenantUserID"
//	@Success		200	{string}	TenantUserID
//	@Failure		500	{string}	TenantUserID
//	@Router			/tenant_users/{id} [delete]
func (c *TenantUserController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.tenantUserSvc.Delete(ctx.Context(), id)
}
