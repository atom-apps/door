package dto

type SendSmsVerifyCodeForm struct {
	Code  string `json:"code" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

type SendEmailVerifyCodeForm struct {
	Address string `json:"phone" validate:"required"`
}
