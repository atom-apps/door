package oauth

import (
	"net/url"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "OAuth"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Auth struct {
	App []*App
}

type App struct {
	Name                 string
	ClientID             string
	ClientSecret         string
	CallbackUrl          string
	TokenDuration        string
	RefreshTokenDuration string
	RegisterFields       []RegisterField
	PasswordComplexRule  *PasswordComplexRule
}

type PasswordComplexRule struct {
	MinLength    *int
	MaxLength    *int
	UpperCase    *int
	SpecialChars *int
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Auth
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	for _, app := range config.App {
		if app.TokenDuration == "" {
			app.TokenDuration = "168h"
		}
		if app.RefreshTokenDuration == "" {
			app.RefreshTokenDuration = "720h"
		}
	}

	return container.Container.Provide(func() *Auth {
		return &config
	}, o.DiOptions()...)
}

func (c *Auth) GetApp(clientID, clientSecret string) *App {
	for _, app := range c.App {
		if app.ClientID == clientID && app.ClientSecret == clientSecret {
			return app
		}
	}
	return nil
}

func (c *Auth) GetAppByName(name string) (*App, error) {
	for _, app := range c.App {
		if app.Name == name {
			return app, nil
		}
	}
	return nil, ErrInvalidApp
}

func (c *App) PasswordComplexRules() []PasswordRuleFunc {
	if c.PasswordComplexRule == nil {
		return []PasswordRuleFunc{
			MinLengthRule(6),
			MaxLengthRule(16),
		}
	}

	rules := []PasswordRuleFunc{}
	if c.PasswordComplexRule.MinLength != nil {
		rules = append(rules, MinLengthRule(*c.PasswordComplexRule.MinLength))
	}

	if c.PasswordComplexRule.MaxLength != nil {
		rules = append(rules, MaxLengthRule(*c.PasswordComplexRule.MaxLength))
	}

	if c.PasswordComplexRule.UpperCase != nil {
		rules = append(rules, UpperCaseRule(*c.PasswordComplexRule.UpperCase))
	}

	if c.PasswordComplexRule.SpecialChars != nil {
		rules = append(rules, SpecialCharRule(*c.PasswordComplexRule.SpecialChars))
	}

	return rules
}

func (c *App) GetCallbackURL(code, scope string) (string, error) {
	u, err := url.Parse(c.CallbackUrl)
	if err != nil {
		return "", err
	}

	query := u.Query()
	query.Set("code", code)
	query.Set("scope", scope)

	u.RawQuery = query.Encode()

	return u.String(), nil
}
