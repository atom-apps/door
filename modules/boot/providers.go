package boot

import (
	"strings"

	"github.com/atom-apps/door/providers/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
	// clientv3 "go.etcd.io/etcd/client/v3"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: provideHttpMiddleware},
		// {Provider: provideGoMicroOptions},
		// {Provider: provideSwagger},
	}
}

// func provideSwagger(opts ...opt.Option) error {
// 	return container.Container.Provide(func(swagger *swagger.Swagger) contracts.Initial {
// 		lo.Must0(swagger.Load(docs.SwaggerSpec))
// 		return nil
// 	}, atom.GroupInitial)
// }

// func provideGoMicroOptions(opts ...opt.Option) error {
// 	_ = container.Container.Provide(func(ctx context.Context, log *log.Logger, client *clientv3.Client) registry.Registry {
// 		logger, _ := zap.NewLogger(
// 			zap.WithLogger(log.Logger),
// 		)

// 		return etcd.NewRegistry(
// 			registry.Logger(logger),
// 			registry.Timeout(time.Second*5),
// 			etcd.Client(client),
// 		)
// 	})

// 	return nil
// }

func provideHttpMiddleware(opts ...opt.Option) error {
	return container.Container.Provide(
		func(httpsvc contracts.HttpService, jwt *jwt.JWT) contracts.Initial {
			engine := httpsvc.GetEngine().(*fiber.App)
			// Initialize default config
			engine.Use(cors.New())
			engine.Static("", "./frontend/dist", fiber.Static{
				Compress: true,
			})
			engine.Use(httpMiddlewareJWT(jwt))
			return nil
		}, atom.GroupInitial)
}

func httpMiddlewareJWT(j *jwt.JWT) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		skipAuth := []string{"/auth", "/v1/auth", "/v1/services"}
		for _, path := range skipAuth {
			if strings.HasPrefix(ctx.Path(), path) {
				return ctx.Next()
			}
		}

		token, ok := ctx.GetReqHeaders()[jwt.HttpHeader]
		if !ok {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		if !strings.HasPrefix(token, jwt.TokenPrefix) {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		token = token[len(jwt.TokenPrefix):]
		claims, err := j.ParseToken(token)
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		ctx.Locals(jwt.CtxKey, claims)

		return ctx.Next()
	}
}
