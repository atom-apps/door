package controller

import (
	"github.com/atom-apps/door/modules/auth/dto"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
)

// @provider
type PermissionController struct {
	casbin *casbin.Casbin
}

// Check
//
//	@Summary		CheckPermission
//	@Description	CheckPermission
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.PermissionCheckForm	true	"PermissionCheckForm"
//	@Router			/v1/auth/permission/check [post]
func (c *PermissionController) Check(ctx *fiber.Ctx, check *dto.PermissionCheckForm) error {
	claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
	if !ok {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	if !c.casbin.Check(claim.UserID, claim.TenantID, check.Path, check.Method) {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return nil
}
