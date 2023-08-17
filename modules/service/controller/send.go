package controller

import (
	authSvc "github.com/atom-apps/door/modules/auth/service"
	"github.com/atom-apps/door/modules/service/dto"
	"github.com/atom-apps/door/modules/service/service"
	userSvc "github.com/atom-apps/door/modules/user/service"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/captcha"
	"github.com/gofiber/fiber/v2"
)

// @provider
type SendController struct {
	userSvc *userSvc.UserService
	authSvc *authSvc.AuthService
	svc     *service.SendService
	captcha *captcha.Captcha
}

// Sms send sms code
//
//	@Summary	send sms code
//	@Tags		Service
//	@Accept		json
//	@Produce	json
//	@Param		body	body	dto.SendSmsVerifyCodeForm	true	"SendSmsVerifyCodeForm"
//	@Router		/services/send/sms [post]
func (c *SendController) Sms(ctx *fiber.Ctx, body *dto.SendVerifyCodeForm) error {
	if !c.userSvc.IsPhoneValid(ctx.Context(), body.To) {
		return oauth.ErrPhoneInvalid
	}

	if !c.captcha.Verify(body.CaptchaID, body.Code) {
		return oauth.ErrVerifyCodeInvalid
	}

	return c.svc.SendSmsCode(ctx.Context(), body.To)
}

// Email send email code
//
//	@Summary	send email code
//	@Tags		Service
//	@Accept		json
//	@Produce	json
//	@Param		body	body	dto.SendEmailVerifyCodeForm	true	"SendEmailVerifyCodeForm"
//	@Router		/services/send/email [post]
func (c *SendController) Email(ctx *fiber.Ctx, body *dto.SendVerifyCodeForm) error {
	if !c.userSvc.IsEmailValid(ctx.Context(), body.To) {
		return oauth.ErrEmailInvalid
	}

	if !c.captcha.Verify(body.CaptchaID, body.Code) {
		return oauth.ErrVerifyCodeInvalid
	}

	return c.svc.SendEmailCode(ctx.Context(), body.To)
}
