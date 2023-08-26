package controller

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	userSvc "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/casbin"
	"github.com/atom-providers/uuid"
	"github.com/gofiber/fiber/v2"
)

// @provider
type PageController struct {
	casbin     *casbin.Casbin
	uuid       *uuid.Generator
	oauth      *oauth.Auth
	userSvc    *userSvc.UserService
	sessionSvc *userSvc.SessionService
	tokenSvc   *userSvc.TokenService
}

func (c *PageController) sendCookie(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:     consts.SessionName,
		Value:    c.uuid.MustGenerate(),
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Secure:   true,
		HTTPOnly: true,
	})
}

func (c *PageController) canAutoLogin(ctx *fiber.Ctx) (string, bool) {
	session := ctx.Cookies(consts.SessionName, "")
	if session == "" {
		return "", false
	}

	sess, err := c.sessionSvc.GetBySessionID(ctx.Context(), session)
	if err != nil {
		return "", false
	}

	user, err := c.userSvc.GetByID(ctx.Context(), sess.UserID)
	if err != nil {
		return "", false
	}

	token, err := c.tokenSvc.GetByUserID(ctx.Context(), user.ID)
	if err != nil {
		return "", false
	}

	redirect, err := c.oauth.GetCallbackURL(token.Code, token.Scope, ctx.Query("redirect", ""))
	if err != nil {
		return "", false
	}

	return redirect, true
}

// Login Login Page
//
//	@Summary		Login Page
//	@Description	Login Page
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/signin [get]
func (c *PageController) Signin(ctx *fiber.Ctx, appName string) error {
	if redirect, can := c.canAutoLogin(ctx); can {
		return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
	}
	c.sendCookie(ctx)
	return ctx.Render("./frontend/dist/index.html", fiber.Map{})
}

// Signup page
//
//	@Summary		Signup page
//	@Description	Signup Page
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/signup [get]
func (c *PageController) Signup(ctx *fiber.Ctx) error {
	if redirect, can := c.canAutoLogin(ctx); can {
		return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
	}
	c.sendCookie(ctx)
	return ctx.Render("./frontend/dist/index.html", fiber.Map{})
}

// ResetPassword page
//
//	@Summary		ResetPassword
//	@Description	ResetPassword
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/reset-password [get]
func (c *PageController) ResetPassword(ctx *fiber.Ctx) error {
	if redirect, can := c.canAutoLogin(ctx); can {
		return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
	}
	c.sendCookie(ctx)
	return ctx.Render("./frontend/dist/index.html", fiber.Map{})
}
