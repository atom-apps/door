package dto

type SendVerifyCodeForm struct {
	CaptchaID string `json:"captcha_id,omitempty"`
	Code      string `json:"code,omitempty" validate:"required"`
	To        string `json:"to,omitempty" validate:"required"`
}
