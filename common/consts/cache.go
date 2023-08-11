package consts

import "fmt"

// swagger:enum CacheKey
// ENUM(
// RegisterCode = "code:register:%s",
// )
type CacheKey string

func (c CacheKey) With(args ...any) string {
	return fmt.Sprintf(string(c), args...)
}
