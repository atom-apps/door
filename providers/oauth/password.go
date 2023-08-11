package oauth

import (
	"fmt"
	"regexp"
)

type PasswordRuleFunc func(string) bool

func MinLengthRule(count int) PasswordRuleFunc {
	return func(password string) bool {
		return len(password) >= count
	}
}

func MaxLengthRule(count int) PasswordRuleFunc {
	return func(password string) bool {
		return len(password) <= count
	}
}

func UpperCaseRule(count int) PasswordRuleFunc {
	return func(password string) bool {
		regex := fmt.Sprintf(`[A-Z]{%d,}`, count)
		match, _ := regexp.MatchString(regex, password)
		return match
	}
}

func DigitRule(count int) PasswordRuleFunc {
	return func(password string) bool {
		regex := fmt.Sprintf(`[0-9]{%d,}`, count)
		match, _ := regexp.MatchString(regex, password)
		return match
	}
}

func SpecialCharRule(count int) PasswordRuleFunc {
	return func(password string) bool {
		regex := fmt.Sprintf(`[!@#$%%^&*()-=_+\\|[\]{};:'",.<>/?]{%d,}`, count)
		match, _ := regexp.MatchString(regex, password)
		return match
	}
}
