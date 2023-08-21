package boot

import (
	"strings"

	"github.com/atom-apps/door/common/consts"
	userModule "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/jwt"
	"github.com/atom-providers/casbin"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func httpMiddlewareCasbin(
	casbin *casbin.Casbin,
	roleSvc *userModule.RoleService,
	tenantSvc *userModule.TenantService,
) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		for _, path := range skipJwt {
			if strings.HasPrefix(ctx.Path(), path) {
				return ctx.Next()
			}
		}

		for _, path := range skipAuth {
			if ctx.Path() == path {
				return ctx.Next()
			}
		}

		claim, ok := ctx.Locals(jwt.CtxKey).(*jwt.Claims)
		if !ok {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		role, err := roleSvc.GetByUserID(ctx.Context(), claim.TenantID, claim.UserID)
		if err != nil {
			return errors.Wrap(err, "middleware: get user role failed")
		}

		if role.Slug == consts.RoleSuperAdmin.String() {
			return ctx.Next()
		}

		if casbin.Check(claim.UserID, claim.TenantID, ctx.Path(), ctx.Method()) {
			return ctx.Next()
		}

		return ctx.SendStatus(fiber.StatusForbidden)
	}
}
