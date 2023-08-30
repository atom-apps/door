package boot

import (
	userModule "github.com/atom-apps/door/modules/users/service"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

var (
	skipJwt = []string{
		"/auth/signin",
		"/auth/signup",
		"/auth/reset-password",

		"/v1/auth/check-reset-password-code",
		"/v1/auth/exchange-token-by-code",
		"/v1/auth/signin",
		"/v1/auth/signup",

		"/v1/services/captcha/generate",
		"/v1/services/send/sms",
		"/v1/services/send/email",
	}
	skipAuth = []string{"/v1/permission/check"}
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideHttpMiddleware},
	}
}

func provideHttpMiddleware(opts ...opt.Option) error {
	return container.Container.Provide(func(
		httpsvc contracts.HttpService,
		jwt *jwt.JWT,
		casbin *casbin.Casbin,
		roleSvc *userModule.RoleService,
		tenantSvc *userModule.TenantService,
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
