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
	Method    oauth.SignInMethod `form:"method" json:"method,omitempty"`
	Username  string             `form:"username" json:"username,omitempty"`
	Code      *string            `form:"code" json:"code,omitempty"`
	Password  *string            `form:"password" json:"password,omitempty"`
	Captcha   *string            `form:"captcha" json:"captcha,omitempty"`
	CaptchaID *string            `form:"captcha_id" json:"captcha_id,omitempty"`
	Token     bool               `json:"token,omitempty"`
}

type LogoutForm struct {
	Token string `json:"token,omitempty"`
}

type RefreshTokenForm struct {
	RefreshToken string `json:"token,omitempty"`
}

type ExchangeTokenByCodeForm struct {
	Code     string `json:"code,omitempty"`
	Scope    string `json:"scope,omitempty"`
	Redirect string `json:"redirect,omitempty"`
	Token    string `json:"token,omitempty"`
}

type CheckPasswordResetCodeForm struct {
	Code     string `json:"code,omitempty"`
	Username string `json:"username,omitempty"`
}

type CheckPasswordResetToken struct {
	Token string `json:"token,omitempty"`
}

type ResetPasswordForm struct {
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty"`
}
