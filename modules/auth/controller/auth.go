package controller

import (
	"time"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/modules/auth/dto"
	"github.com/atom-apps/door/modules/auth/service"
	userSvc "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"golang.org/x/oauth2"
)

// @provider
type AuthController struct {
	oauth      *oauth.Auth
	authSvc    *service.AuthService
	userSvc    *userSvc.UserService
	sessionSvc *userSvc.SessionService
	tokenSvc   *userSvc.TokenService
}

func (c *AuthController) canAutoLogin(ctx *fiber.Ctx, appName string) (string, bool) {
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
func (c *AuthController) SigninPage(ctx *fiber.Ctx, appName string) error {
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
func (c *AuthController) SignupPage(ctx *fiber.Ctx, appName string) error {
	return ctx.Render("./frontend/dist/index.html", fiber.Map{})
}

// Signup
//
//	@Summary		Signup
//	@Description	Signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.SignUpForm	true	"SignUpForm"
//	@Success			200	{object}	dto.ExchangeTokenByCodeForm
//	@Router			/auth/signup [post]
func (c *AuthController) SignUp(ctx *fiber.Ctx, form *dto.SignUpForm) (*dto.ExchangeTokenByCodeForm, error) {
	app, err := c.oauth.GetAppByName(form.AppName)
	if err != nil {
		return nil, err
	}

	if err := c.authSvc.SignUpCheckRegisterMethod(ctx.Context(), form, app); err != nil {
		return nil, err
	}

	if _, err = c.authSvc.CreateUser(ctx.Context(), form); err != nil {
		return nil, err
	}

	return c.SignIn(ctx, &dto.SignInForm{
		AppName:  form.AppName,
		SID:      form.SID,
		Username: *common.OneOf(form.Username, form.Email, form.Phone),
		Password: form.Password,
		Method:   oauth.SignInMethodPassword,
	})
}

// SignIn Signin
//
//	@Summary		Signin
//	@Description	Signin
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.SignInForm	true	"SignInForm"
//	@Success			200	{object}	dto.ExchangeTokenByCodeForm
//	@Router			/auth/signin [post]
func (c *AuthController) SignIn(ctx *fiber.Ctx, form *dto.SignInForm) (*dto.ExchangeTokenByCodeForm, error) {
	app, err := c.oauth.GetAppByName(form.AppName)
	if err != nil {
		return nil, err
	}

	user, err := c.userSvc.GetByUsernameOrEmailOrPhone(ctx.Context(), form.Username)
	if err != nil {
		return nil, err
	}

	if err := c.authSvc.VerifySignInPasswordOrCode(ctx.Context(), form, user); err != nil {
		return nil, err
	}

	// write user session id
	sess, err := c.sessionSvc.CreateForUser(ctx.Context(), user.ID, form.SID)
	if err != nil {
		return nil, err
	}

	token, err := c.tokenSvc.CreateForUser(ctx.Context(), user.ID, sess.ID, app)
	if err != nil {
		return nil, err
	}

	redirect, err := app.GetCallbackURL(token.Code, token.Scope, ctx.Query("redirect", ""))
	if err != nil {
		return nil, err
	}

	return &dto.ExchangeTokenByCodeForm{
		Code:     token.Code,
		Scope:    token.Scope,
		Redirect: redirect,
	}, nil
}

// Logout
//
//	@Summary		Logout
//	@Description	Logout
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	TODO:AddData
//	@Router			/auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx, form *dto.LogoutForm) error {
	_, err := c.oauth.GetAppByName(form.AppName)
	if err != nil {
		return err
	}

	token, err := c.tokenSvc.GetByToken(ctx.Context(), form.Token, form.AppName)
	if err != nil {
		return err
	}

	return c.sessionSvc.DeleteBySessionID(ctx.Context(), token.SessionID)
}

// RefreshToken RefreshToken
//
//	@Summary		RefreshToken
//	@Description	RefreshToken
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	TODO:AddData
//	@Router			/auth/refresh-token [post]
func (c *AuthController) RefreshToken(ctx *fiber.Ctx, form *dto.RefreshTokenForm) (*oauth2.Token, error) {
	app, err := c.oauth.GetAppByName(form.AppName)
	if err != nil {
		return nil, err
	}

	token, err := c.tokenSvc.GetByRefreshToken(ctx.Context(), form.RefreshToken, form.AppName)
	if err != nil {
		return nil, err
	}

	token, err = c.tokenSvc.RefreshToken(ctx.Context(), token, app)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.CreatedAt.Add(lo.Must(time.ParseDuration(app.TokenDuration))).Add(-time.Minute * 10),
	}, nil
}

// GetTokenByCode GetTokenByCode
//
//	@Summary		GetTokenByCode
//	@Description	GetTokenByCode
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	oauth2.Token
//	@Router			/auth/exchange-token-by-code [post]
func (c *AuthController) ExchangeTokenByCode(ctx *fiber.Ctx, form *dto.ExchangeTokenByCodeForm) (*oauth2.Token, error) {
	return c.tokenSvc.GetOAuthTokenByCode(ctx.Context(), form.Code, form.Scope)
}
