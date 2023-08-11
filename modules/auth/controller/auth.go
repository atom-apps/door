package controller

import (
	"net/http"
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

	redirect, err := app.GetSignInRedirectURL(token.Code, token.Scope)
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
	if u, ok := c.canAutoLogin(ctx, appName); ok {
		return ctx.Redirect(u, http.StatusTemporaryRedirect)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     consts.SessionName,
		Value:    c.sessionSvc.GenerateSessionID(),
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HTTPOnly: true,
	})
	// todo: render page
	return nil
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
	if u, ok := c.canAutoLogin(ctx, appName); ok {
		return ctx.Redirect(u, http.StatusTemporaryRedirect)
	}

	// todo: render signup page
	return nil
}

// Signup
//
//	@Summary		Signup
//	@Description	Signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.SignUpForm	true	"SignUpForm"
//	@Router			/auth/signup/{app_name} [post]
func (c *AuthController) SignUp(ctx *fiber.Ctx, appName string, form *dto.SignUpForm) error {
	app, err := c.oauth.GetAppByName(appName)
	if err != nil {
		return err
	}

	session := ctx.Cookies(consts.SessionName, "")
	if session == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := c.authSvc.SignUpCheckRegisterMethod(ctx.Context(), form, app); err != nil {
		return err
	}

	if _, err = c.authSvc.CreateUser(ctx.Context(), form); err != nil {
		return err
	}

	return c.SignIn(ctx, appName, &dto.SignInForm{
		Username: common.OneOf(form.Username, form.Email, form.Phone),
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
//	@Param			body	body		dto.SignInForm	true	"SignInForm"
//	@Router			/auth/signin/{app_name} [post]
func (c *AuthController) SignIn(ctx *fiber.Ctx, appName string, form *dto.SignInForm) error {
	app, err := c.oauth.GetAppByName(appName)
	if err != nil {
		return err
	}

	session := ctx.Cookies(consts.SessionName, "")
	if session == "" {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user, err := c.userSvc.GetByUsernameOrEmailOrPhone(ctx.Context(), *form.Username)
	if err != nil {
		return err
	}

	if err := c.authSvc.VerifySignInPasswordOrCode(ctx.Context(), form, user); err != nil {
		return err
	}

	// write user session id
	sess, err := c.sessionSvc.CreateForUser(ctx.Context(), user.ID, session)
	if err != nil {
		return err
	}

	token, err := c.tokenSvc.CreateForUser(ctx.Context(), user.ID, sess.ID, app)
	if err != nil {
		return err
	}

	redirect, err := app.GetSignInRedirectURL(token.Code, token.Scope)
	if err != nil {
		return err
	}

	return ctx.Redirect(redirect, http.StatusTemporaryRedirect)
}

// Logout
//
//	@Summary		Logout
//	@Description	Logout
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200			{string}	TODO:AddData
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
//	@Success		200			{string}	TODO:AddData
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
