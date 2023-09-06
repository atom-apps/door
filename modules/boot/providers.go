package boot

import (
	"context"

	userSvc "github.com/atom-apps/door/modules/users/service"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

var skipJwt = []string{
	"/auth/signin",
	"/auth/signup",
	"/auth/reset-password",

	"/v1/auth/check-reset-password-code",
	"/v1/auth/exchange-token-by-code",
	"/v1/auth/signin",
	"/v1/auth/signup",
	"/v1/auth/test",

	"/v1/services/captcha/generate",
	"/v1/services/send/sms",
	"/v1/services/send/email",
}

func Providers() container.Providers {
	return container.Providers{
		{Provider: providePermissionRules},
		{Provider: provideHttpMiddleware},
	}
}

func provideHttpMiddleware(opts ...opt.Option) error {
	return container.Container.Provide(func(
		httpsvc contracts.HttpService,
		jwt *jwt.JWT,
		casbin *casbin.Casbin,
		roleSvc *userSvc.RoleService,
		tenantSvc *userSvc.TenantService,
	) contracts.Initial {
		engine := httpsvc.GetEngine().(*fiber.App)
		// Initialize default config
		engine.Use(cors.New())
		engine.Static("", "./frontend/dist", fiber.Static{
			Compress: true,
		})
		engine.Use(httpMiddlewareJWT(jwt))
		engine.Use(httpMiddlewareCasbin(casbin, roleSvc, tenantSvc))
		return nil
	}, atom.GroupInitial)
}

func providePermissionRules(opts ...opt.Option) error {
	return container.Container.Provide(func(
		permissionSvc *userSvc.PermissionService,
		userTenantRoleSvc *userSvc.UserTenantRoleService,
		casbin *casbin.Casbin,
	) contracts.Initial {
		groups, err := userTenantRoleSvc.CasbinGroups(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if _, err := casbin.LoadGroups(groups); err != nil {
			log.Fatal(err)
		}
		log.Infof("load groups: %d", len(groups))

		// permissions
		policies, err := permissionSvc.CasbinPolicies(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if _, err := casbin.LoadPolicies(policies); err != nil {
			log.Fatal(err)
		}
		log.Infof("load policies: %d", len(policies))

		return nil
	}, atom.GroupInitial)
}
