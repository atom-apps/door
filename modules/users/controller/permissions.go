package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/users/service"
	"github.com/gofiber/fiber/v2"
)

// @provider
type PermissionRuleController struct {
	permissionRuleSvc *service.PermissionRuleService
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
func (c *PermissionRuleController) AttachUsers(ctx *fiber.Ctx, id, tenantID int64, users *common.IDsForm) error {
	return c.permissionRuleSvc.AddRoleUsers(ctx.Context(), tenantID, id, users.IDs)
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
//	@Router		/v1/permissions/detach/{role_id}/{tenant_id} [put]
func (c *PermissionRuleController) DetachUsers(ctx *fiber.Ctx, id, tenantID int64, users *common.IDsForm) error {
	return c.permissionRuleSvc.DeleteRoleUsers(ctx.Context(), tenantID, id, users.IDs)
}
