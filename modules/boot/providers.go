package boot

import (
	"context"
	"strconv"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	userModule "github.com/atom-apps/door/modules/users/service"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/samber/lo"
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
		{Provider: providePermissionRules},
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

func providePermissionRules(opts ...opt.Option) error {
	return container.Container.Provide(func(query *query.Query, casbin *casbin.Casbin) contracts.Initial {
		// init groups
		groupModels, err := query.UserTenantRole.WithContext(context.Background()).Find()
		if err != nil {
			log.Fatal(err)
		}
		groups := lo.Map(groupModels, func(item *models.UserTenantRole, _ int) []string {
			return []string{
				strconv.Itoa(int(item.UserID)),
				strconv.Itoa(int(item.TenantID)),
				strconv.Itoa(int(item.RoleID)),
			}
		})

		if _, err := casbin.LoadGroups(groups); err != nil {
			log.Fatal(err)
		}
		log.Infof("load groups: %d", len(groups))

		// permissions
		permissionModels, err := query.Permission.WithContext(context.Background()).Find()
		if err != nil {
			log.Fatal(err)
		}
		// TODO: cal permissions
		permissions := lo.Map(permissionModels, func(item *models.Permission, _ int) []string {
			return []string{
				strconv.Itoa(int(item.RoleID)),
				strconv.Itoa(int(item.TenantID)),
				"/*",
				"GET",
			}
		})
		if _, err := casbin.LoadPolicies(permissions); err != nil {
			log.Fatal(err)
		}
		log.Infof("load policies: %d", len(groups))

		return nil
	}, atom.GroupInitial)
}
