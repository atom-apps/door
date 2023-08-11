package dto

import "github.com/atom-apps/door/providers/oauth"

type SignUpForm struct {
	Username *string `json:"username,omitempty"`

	Email     *string `json:"email,omitempty"`
	EmailCode *string `json:"email_code,omitempty"`

	Phone     *string `json:"phone,omitempty"`
	PhoneCode *string `json:"phone_code,omitempty"`

	Password *string `json:"password,omitempty"`

	Captcha   *string `json:"captcha,omitempty"`
	CaptchaID *string `json:"captcha_id,omitempty"`
}

type SignInForm struct {
	Method    oauth.SignInMethod `json:"method"`
	Username  *string            `json:"username,omitempty"`
	Code      *string            `json:"code,omitempty"`
	Password  *string            `json:"password,omitempty"`
	Captcha   *string            `json:"captcha,omitempty"`
	CaptchaID *string            `json:"captcha_id,omitempty"`
}

type LogoutForm struct {
	Token   string `json:"token,omitempty"`
	AppName string `json:"app_name,omitempty"`
}

type RefreshTokenForm struct {
	RefreshToken string `json:"token,omitempty"`
	AppName      string `json:"app_name,omitempty"`
}
