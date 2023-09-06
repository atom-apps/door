package boot

import (
	userModule "github.com/atom-apps/door/modules/users/service"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
)

func httpMiddlewareCasbin(
	casbin *casbin.Casbin,
	roleSvc *userModule.RoleService,
	tenantSvc *userModule.TenantService,
) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
		if !ok {
			return ctx.Next()
		}

		if claim.Role == jwt.RoleSuperAdmin.String() {
			return ctx.Next()
		}

		if casbin.Check(claim.UserID, claim.TenantID, ctx.Path(), ctx.Method()) {
			return ctx.Next()
		}

		return ctx.SendStatus(fiber.StatusForbidden)
	}
}
