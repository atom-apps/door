package controller

import (
	"time"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/common/errorx"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/auth/dto"
	"github.com/atom-apps/door/modules/auth/service"
	serviceSvc "github.com/atom-apps/door/modules/service/service"
	userSvc "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"golang.org/x/oauth2"
)

// @provider
type AuthController struct {
	oauth         *oauth.Auth
	authSvc       *service.AuthService
	userSvc       *userSvc.UserService
	sessionSvc    *userSvc.SessionService
	tenantSvc     *userSvc.TenantService
	tokenSvc      *userSvc.TokenService
	permissionSvc *userSvc.PermissionRuleService
	sendSvc       *serviceSvc.SendService
}

// Signup
//
//	@Summary		Signup
//	@Description	Signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.SignUpForm	true	"SignUpForm"
//	@Success		200		{object}	dto.ExchangeTokenByCodeForm
//	@Router			/v1/auth/signup [post]
func (c *AuthController) SignUp(ctx *fiber.Ctx, form *dto.SignUpForm) (*dto.ExchangeTokenByCodeForm, error) {
	if err := c.authSvc.SignUpCheckRegisterMethod(ctx.Context(), form); err != nil {
		return nil, err
	}

	if _, err := c.authSvc.CreateUser(ctx.Context(), form); err != nil {
		return nil, err
	}

	return c.SignIn(ctx, &dto.SignInForm{
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
//	@Param			body	body		dto.SignInForm	true	"SignInForm"
//	@Success		200		{object}	dto.ExchangeTokenByCodeForm
//	@Router			/v1/auth/signin [post]
func (c *AuthController) SignIn(ctx *fiber.Ctx, form *dto.SignInForm) (*dto.ExchangeTokenByCodeForm, error) {
	sessID := ctx.Cookies(consts.SessionName, "")
	// if sessID == "" {
	// 	return nil, errorx.ErrInvalidRequest
	// }

	user, err := c.userSvc.GetByUsernameOrEmailOrPhone(ctx.Context(), form.Username)
	if err != nil {
		return nil, err
	}

	if err := c.authSvc.VerifySignInPasswordOrCode(ctx.Context(), form, user); err != nil {
		return nil, err
	}

	// write user session id
	sess, err := c.sessionSvc.CreateForUser(ctx.Context(), user.ID, sessID)
	if err != nil {
		return nil, err
	}

	tenants, err := c.permissionSvc.GetTenantsByUserID(ctx.Context(), user.ID)
	if err != nil {
		return nil, err
	}

	if len(tenants) == 0 {
		return nil, errorx.ErrForbidden
	}
	tenantID := tenants[0].ID

	token, err := c.tokenSvc.CreateForUser(ctx.Context(), user.ID, tenantID, sess.ID)
	if err != nil {
		return nil, err
	}

	if form.Token {
		oauth, err := c.tokenSvc.GetOAuthTokenByCode(ctx.Context(), token.Code)
		if err != nil {
			return nil, err
		}
		return &dto.ExchangeTokenByCodeForm{Token: oauth.AccessToken}, nil
	}

	redirect, err := c.oauth.GetCallbackURL(token.Code, token.Scope, ctx.Query("redirect", ""))
	if err != nil {
		return nil, err
	}

	return &dto.ExchangeTokenByCodeForm{
		Code:     token.Code,
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
//	@Router			/v1/auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx, claim *jwt.Claims) error {
	token, err := c.tokenSvc.GetByUserID(ctx.Context(), claim.UserID)
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
//	@Success		200		{string}	oauth2.Token
//	@Router			/v1/auth/refresh-token [post]
func (c *AuthController) RefreshToken(ctx *fiber.Ctx, claim *jwt.Claims) (*oauth2.Token, error) {
	token, err := c.tokenSvc.GetByUserID(ctx.Context(), claim.UserID)
	if err != nil {
		return nil, err
	}

	token, err = c.tokenSvc.RefreshToken(ctx.Context(), token)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.CreatedAt.Add(lo.Must(time.ParseDuration(c.oauth.TokenDuration))).Add(-time.Minute * 10),
	}, nil
}

// GetTokenByCode GetTokenByCode
//
//	@Summary		GetTokenByCode
//	@Description	GetTokenByCode
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.ExchangeTokenByCodeForm	true	"ExchangeTokenByCodeForm"
//	@Success		200		{object}	oauth2.Token
//	@Router			/v1/auth/exchange-token-by-code [post]
func (c *AuthController) ExchangeTokenByCode(ctx *fiber.Ctx, form *dto.ExchangeTokenByCodeForm) (*oauth2.Token, error) {
	return c.tokenSvc.GetOAuthTokenByCode(ctx.Context(), form.Code)
}

// CheckResetPasswordCoe
//
//	@Summary		CheckResetPasswordCoe
//	@Description	CheckResetPasswordCoe
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.CheckPasswordResetCodeForm	true	"CheckPasswordResetCode"
//	@Success		200		{object}	oauth2.Token
//	@Router			/v1/auth/check-reset-password-code [post]
func (c *AuthController) CheckResetPasswordCoe(ctx *fiber.Ctx, form *dto.CheckPasswordResetCodeForm) (*dto.CheckPasswordResetToken, error) {
	if !c.sendSvc.VerifyCode(ctx.Context(), consts.VerifyCodeChannelResetPassword, form.Username, form.Code) {
		return nil, errorx.ErrInvalidVerifyCode
	}

	var user *models.User
	var err error

	if c.userSvc.IsEmailValid(ctx.Context(), form.Username) {
		user, err = c.userSvc.GetByEmail(ctx.Context(), form.Username)
		if err != nil {
			return nil, err
		}
	} else if c.userSvc.IsPhoneValid(ctx.Context(), form.Username) {
		user, err = c.userSvc.GetByPhone(ctx.Context(), form.Username)
		if err != nil {
			return nil, err
		}
	}
	token, err := c.userSvc.GetUserIDHashToken(ctx.Context(), user)
	if err != nil {
		return nil, err
	}

	return &dto.CheckPasswordResetToken{Token: token}, nil
}

// ResetPassword
//
//	@Summary		ResetPassword
//	@Description	ResetPassword
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.ResetPasswordForm	true	"ResetPassword"
//	@Success		200		{object}	oauth2.Token
//	@Router			/v1/auth/reset-password-by-token [post]
func (c *AuthController) ResetPassword(ctx *fiber.Ctx, form *dto.ResetPasswordForm) error {
	user, err := c.userSvc.GetUserFromHashToken(ctx.Context(), form.Token)
	if err != nil {
		return err
	}

	if err := c.authSvc.CheckPasswordComplex(ctx.Context(), form.Password, c.oauth.PasswordComplexRules()); err != nil {
		return err
	}

	return c.userSvc.ResetPassword(ctx.Context(), user, form.Password)
}
