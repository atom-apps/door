package boot

import (
	userModule "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/jwt"
	"github.com/atom-providers/casbin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
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
