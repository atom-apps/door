package dto

import "github.com/atom-apps/door/common/consts"

type SendVerifyCodeForm struct {
	Channel   consts.VerifyCodeChannel `json:"channel,omitempty"`
	CaptchaID string                   `json:"captcha_id,omitempty"`
	Code      string                   `json:"code,omitempty" validate:"required"`
	To        string                   `json:"to,omitempty" validate:"required"`
}
