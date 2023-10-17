package consts

import (
	"github.com/atom-apps/door/common/ds"
	"github.com/samber/lo"
)

// swagger:enum UserStatus
// ENUM(default="", blocked)
type UserStatus string

func (u UserStatus) Cn() string {
	switch u {
	case UserStatusDefault:
		return "正常"
	case UserStatusBlocked:
		return "禁用"
	default:
		return "-"
	}
}

func UserStatusLabel(withDefault bool) []ds.LabelItem {
	values := UserStatusValues()
	if withDefault {
		values = append([]UserStatus{""}, values...)
	}

	return lo.Map(values, func(v UserStatus, _ int) ds.LabelItem {
		return ds.LabelItem{
			Label: v.Cn(),
			Value: string(v),
		}
	})
}

const SessionName = "sessionid"
