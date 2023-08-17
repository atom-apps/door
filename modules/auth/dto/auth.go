package dto

import "github.com/atom-apps/door/providers/oauth"

type SignUpForm struct {
	AppName string `json:"app_name"`
	SID     string `form:"sid"`

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
	AppName   string             `form:"app_name" json:"app_name,omitempty"`
	SID       string             `form:"sid" json:"sid,omitempty"`
	Method    oauth.SignInMethod `form:"method" json:"method,omitempty"`
	Username  string             `form:"username" json:"username,omitempty"`
	Code      *string            `form:"code" json:"code,omitempty"`
	Password  *string            `form:"password" json:"password,omitempty"`
	Captcha   *string            `form:"captcha" json:"captcha,omitempty"`
	CaptchaID *string            `form:"captcha_id" json:"captcha_id,omitempty"`
}

type LogoutForm struct {
	Token   string `json:"token,omitempty"`
	AppName string `json:"app_name,omitempty"`
}

type RefreshTokenForm struct {
	RefreshToken string `json:"token,omitempty"`
	AppName      string `json:"app_name,omitempty"`
}

type ExchangeTokenByCodeForm struct {
	Code     string `json:"code,omitempty"`
	Scope    string `json:"scope,omitempty"`
	Redirect string `json:"redirect,omitempty"`
}
