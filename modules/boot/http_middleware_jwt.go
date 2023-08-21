package boot

import (
	"strings"

	"github.com/atom-apps/door/providers/jwt"
	"github.com/gofiber/fiber/v2"
)

func httpMiddlewareJWT(j *jwt.JWT) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		for _, path := range skipJwt {
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
