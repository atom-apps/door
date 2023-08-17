package oauth

import (
	"fmt"
	"regexp"
)

type PasswordRuleFunc func(string) error

func MinLengthRule(count int) PasswordRuleFunc {
	return func(password string) error {
		if len(password) <= count {
			return fmt.Errorf("最少长度为 %d 位字符", count)
		}
		return nil
	}
}

func MaxLengthRule(count int) PasswordRuleFunc {
	return func(password string) error {
		if len(password) >= count {
			return fmt.Errorf("最大长度为 %d 位字符", count)
		}
		return nil
	}
}

func UpperCaseRule(count int) PasswordRuleFunc {
	return func(password string) error {
		regex := fmt.Sprintf(`[A-Z]{%d,}`, count)
		match, _ := regexp.MatchString(regex, password)
		if !match {
			return fmt.Errorf("最小需包含 %d 个大写字符", count)
		}
		return nil
	}
}

func DigitRule(count int) PasswordRuleFunc {
	return func(password string) error {
		regex := fmt.Sprintf(`[0-9]{%d,}`, count)
		match, _ := regexp.MatchString(regex, password)
		if !match {
			return fmt.Errorf("最小需包含 %d 个数字字符", count)
		}
		return nil
	}
}

func SpecialCharRule(count int) PasswordRuleFunc {
	return func(password string) error {
		regex := fmt.Sprintf(`[!@#$%%^&*()-=_+\\|[\]{};:'",.<>/?]{%d,}`, count)
		match, _ := regexp.MatchString(regex, password)
		if !match {
			return fmt.Errorf("最小需包含 %d 个特殊字符", count)
		}
		return nil
	}
}
