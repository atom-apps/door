package controller

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	userSvc "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/uuid"
	"github.com/gofiber/fiber/v2"
)

// @provider
type PageController struct {
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

func (c *PageController) canAutoLogin(ctx *fiber.Ctx, appName string) (string, bool) {
	app, err := c.oauth.GetAppByName(appName)
	if err != nil {
		return "", false
	}

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

	token, err := c.tokenSvc.GetByUserID(ctx.Context(), user.ID, app.Name)
	if err != nil {
		return "", false
	}

	redirect, err := app.GetCallbackURL(token.Code, token.Scope, ctx.Query("redirect", ""))
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
//	@Router			/auth/signin/{app_name} [get]
func (c *PageController) Signin(ctx *fiber.Ctx, appName string) error {
	if redirect, can := c.canAutoLogin(ctx, appName); can {
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
//	@Router			/auth/signup/{app_name} [get]
func (c *PageController) Signup(ctx *fiber.Ctx, appName string) error {
	if redirect, can := c.canAutoLogin(ctx, appName); can {
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
//	@Router			/auth/reset-password/{app_name} [get]
func (c *PageController) ResetPassword(ctx *fiber.Ctx, appName string) error {
	if redirect, can := c.canAutoLogin(ctx, appName); can {
		return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
	}
	c.sendCookie(ctx)
	return ctx.Render("./frontend/dist/index.html", fiber.Map{})
}
