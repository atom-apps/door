package controller

import (
	"github.com/atom-apps/door/modules/service/dto"
	"github.com/atom-providers/captcha"
	"github.com/gofiber/fiber/v2"
)

// @provider
type CaptchaController struct {
	captcha *captcha.Captcha
}

// Generate new captcha
//
//	@Summary	captcha
//	@Tags		Service
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.Captcha
//	@Router		/services/captcha/generate [get]
func (c *CaptchaController) Generate(ctx *fiber.Ctx) (*dto.Captcha, error) {
	image, err := c.captcha.Generate()
	if err != nil {
		return nil, err
	}
	return &dto.Captcha{
		Image: image.PicPath,
		ID:    image.CaptchaId,
	}, nil
}
