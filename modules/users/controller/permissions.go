package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	systemDto "github.com/atom-apps/door/modules/systems/dto"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type PermissionController struct {
	permissionSvc     *service.PermissionService
	userTenantRoleSvc *service.UserTenantRoleService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"PermissionID"
//	@Success		200	{object}	dto.PermissionItem
//	@Router			/v1/users/permissions/{id} [get]
func (c *PermissionController) Show(ctx *fiber.Ctx, id uint64) (*dto.PermissionItem, error) {
	item, err := c.permissionSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.permissionSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.PermissionListQueryFilter	true	"PermissionListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.PermissionItem}
//	@Router			/v1/users/permissions [get]
func (c *PermissionController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.PermissionListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.permissionSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.permissionSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.PermissionForm	true	"PermissionForm"
//	@Success		200		{string}	PermissionID
//	@Router			/v1/users/permissions [post]
func (c *PermissionController) Create(ctx *fiber.Ctx, body *dto.PermissionForm) error {
	return c.permissionSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"PermissionID"
//	@Param			body	body		dto.PermissionForm	true	"PermissionForm"
//	@Success		200		{string}	PermissionID
//	@Router			/v1/users/permissions/{id} [put]
func (c *PermissionController) Update(ctx *fiber.Ctx, id uint64, body *dto.PermissionForm) error {
	return c.permissionSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"PermissionID"
//	@Success		200	{string}	PermissionID
//	@Router			/v1/users/permissions/{id} [delete]
func (c *PermissionController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.permissionSvc.Delete(ctx.Context(), id)
}

// Attach users
//
//	@Summary	attach users
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		role_id		path		int				true	"RoleID"
//	@Param		tenant_id	path		int				true	"TenantID"
//	@Param		body		body		common.IDsForm	true	"IDsForm"
//	@Success	200			{string}	RoleID
//	@Failure	500			{string}	RoleID
//	@Router		/v1/users/permissions/attach/{role_id}/{tenant_id} [put]
func (c *PermissionController) AttachUsers(ctx *fiber.Ctx, roleID, tenantID uint64, users *common.IDsForm) error {
	return c.userTenantRoleSvc.AttachUsers(ctx.Context(), tenantID, roleID, users.IDs)
}

// Detach users
//
//	@Summary	detach users
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		id			path		int				true	"RoleID"
//	@Param		tenant_id	path		int				true	"TenantID"
//	@Param		body		body		common.IDsForm	true	"IDsForm"
//	@Success	200			{string}	RoleID
//	@Failure	500			{string}	RoleID
//	@Router		/v1/users/permissions/detach/{role_id}/{tenant_id} [put]
func (c *PermissionController) DetachUsers(ctx *fiber.Ctx, roleID, tenantID uint64, users *common.IDsForm) error {
	return c.userTenantRoleSvc.DetachUsers(ctx.Context(), tenantID, roleID, users.IDs)
}

// Tree
//
//	@Summary	Tree
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Success	200			{string}	RoleID
//	@Router		/v1/users/permissions/tree [get]
func (c *PermissionController) Tree(ctx *fiber.Ctx) ([]systemDto.RouteItem, error) {
	return []systemDto.RouteItem{
		{
			ID:       0,
			Type:     consts.RouteTypeApi,
			ParentID: 0,
			Name:     "",
			Path:     "",
			Metadata: common.RouteMetadata{
				RequiresAuth: false,
				Icon:         "",
				Order:        0,
			},
			Children: []*systemDto.RouteItem{},
		},
	}, nil
}
