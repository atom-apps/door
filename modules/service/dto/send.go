package dto

type SendSmsVerifyCodeForm struct {
	CaptchaID string `json:"captcha_id,omitempty"`
	Code      string `json:"code,omitempty" validate:"required"`
	Phone     string `json:"phone,omitempty" validate:"required"`
}

type SendEmailVerifyCodeForm struct {
	Address string `json:"phone" validate:"required"`
}
