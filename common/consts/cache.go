package consts

import (
	"fmt"
	"strings"
)

// swagger:enum CacheKey
// ENUM(
// VerifyCode = "code:__CHANNEL__:%s",
// )
type CacheKey string

// swagger:enum VerifyCodeChannel
// ENUM(signin, signup, reset-password)
type VerifyCodeChannel string

var AuthChannels = []VerifyCodeChannel{VerifyCodeChannelResetPassword}

func (c CacheKey) VerifyCode(channel VerifyCodeChannel, args ...any) string {
	return fmt.Sprintf(strings.ReplaceAll(c.String(), "__CHANNEL__", channel.String()), args...)
}
