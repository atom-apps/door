package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/common/errorx"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/auth/dto"
	serviceSvc "github.com/atom-apps/door/modules/service/service"
	"github.com/atom-apps/door/modules/users/service"
	"github.com/atom-apps/door/providers/bcrypt"
	"github.com/atom-apps/door/providers/oauth"
	"github.com/atom-providers/uuid"
)

// @provider
type AuthService struct {
	uuid    *uuid.Generator
	hash    *bcrypt.Hash
	userSvc *service.UserService
	sendSvc *serviceSvc.SendService
	auth    *oauth.Auth
}

func (svc *AuthService) SignUpCheckRegisterMethod(ctx context.Context, form *dto.SignUpForm) error {
	for _, m := range svc.auth.RegisterFields {
		switch m {
		case oauth.RegisterFieldEmail:
			if form.Email == nil {
				return oauth.ErrEmailRequired
			}

			if !svc.userSvc.IsEmailValid(ctx, *form.Email) {
				return oauth.ErrEmailInvalid
			}

			if user, _ := svc.userSvc.GetByEmail(ctx, *form.Email); user != nil {
				return oauth.ErrEmailExists
			}

			if !svc.sendSvc.VerifyCode(ctx, consts.VerifyCodeChannelSignup, *form.Email, *form.EmailCode) {
				return oauth.ErrVerifyCodeInvalid
			}

		case oauth.RegisterFieldPhone:
			if form.Phone == nil {
				return oauth.ErrPhoneRequired
			}

			if ok := svc.userSvc.IsPhoneValid(ctx, *form.Phone); !ok {
				return oauth.ErrPhoneInvalid
			}

			if user, _ := svc.userSvc.GetByPhone(ctx, *form.Phone); user != nil {
				return oauth.ErrPhoneExists
			}

			if !svc.sendSvc.VerifyCode(ctx, consts.VerifyCodeChannelSignup, *form.Phone, *form.PhoneCode) {
				return oauth.ErrVerifyCodeInvalid
			}
		case oauth.RegisterFieldUsername:
			if form.Username == nil {
				return oauth.ErrUsernameRequired
			}
			if ok := svc.userSvc.IsUsernameValid(ctx, *form.Username); !ok {
				return oauth.ErrUsernameInvalid
			}
		case oauth.RegisterFieldPassword:
			if form.Password == nil {
				return oauth.ErrPasswordRequired
			}
			if err := svc.CheckPasswordComplex(ctx, *form.Password, svc.auth.PasswordComplexRules()); err != nil {
				return oauth.ErrUsernameInvalid
			}
		}
	}

	return nil
}

func (svc *AuthService) CheckPasswordComplex(ctx context.Context, password string, method []oauth.PasswordRuleFunc) error {
	for _, m := range method {
		if err := m(password); err != nil {
			return err
		}
	}
	return nil
}

// CreateUser
func (svc *AuthService) CreateUser(ctx context.Context, form *dto.SignUpForm) (*models.User, error) {
	m, _ := svc.userSvc.GetByUsernameOrEmailOrPhone(ctx, *common.OneOf(form.Username, form.Email, form.Phone))
	if m != nil {
		return nil, errorx.ErrorUsernameOrEmailOrPhoneAlreadyExists
	}

	model := &models.User{
		UUID:          svc.uuid.MustGenerate(),
		Username:      common.PtrToValue(form.Username, ""),
		Password:      svc.hash.Hash(common.PtrToValue(form.Password, "")),
		Email:         common.PtrToValue(form.Email, ""),
		Phone:         common.PtrToValue(form.Phone, ""),
		EmailVerified: true,
		DisplayName:   "",
		Avatar:        "",
	}

	if err := svc.userSvc.CreateFromModel(ctx, model); err != nil {
		return nil, err
	}

	return model, nil
}

// ComparePassword
func (svc *AuthService) ComparePassword(ctx context.Context, user *models.User, password string) error {
	if svc.hash.Compare(password, user.Password) {
		return nil
	}
	return errorx.ErrorUsernameOrPasswordInvalid
}

// VerifySignInPasswordOrCode
func (svc *AuthService) VerifySignInPasswordOrCode(ctx context.Context, form *dto.SignInForm, user *models.User) error {
	if form.Method == oauth.SignInMethodCode {
		if !svc.sendSvc.VerifyCode(ctx, consts.VerifyCodeChannelSignin, form.Username, *form.Code) {
			return oauth.ErrVerifyCodeInvalid
		}
	}

	if form.Method == oauth.SignInMethodPassword {
		if err := svc.ComparePassword(ctx, user, *form.Password); err != nil {
			return err
		}
	}
	return nil
}
