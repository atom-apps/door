package dto

import (
	"time"

	"github.com/atom-apps/door/common/consts"
)

type UserInfoForm struct {
	UserID      int64         `form:"user_id" json:"user_id,omitempty"`           //
	Affiliation string        `form:"affiliation" json:"affiliation,omitempty"`   // 工作单位
	Title       string        `form:"title" json:"title,omitempty"`               // 职称
	IDCardType  string        `form:"id_card_type" json:"id_card_type,omitempty"` // 证件类型
	IDCard      string        `form:"id_card" json:"id_card,omitempty"`           //
	Biography   string        `form:"biography" json:"biography,omitempty"`       // 自我介绍
	Tag         string        `form:"tag" json:"tag,omitempty"`                   //
	Language    string        `form:"language" json:"language,omitempty"`         //
	Gender      consts.Gender `form:"gender" json:"gender,omitempty"`             // 性别
	Birthday    time.Time     `form:"birthday" json:"birthday,omitempty"`         // 生日
	Education   string        `form:"education" json:"education,omitempty"`       // 学历
	RealName    string        `form:"real_name" json:"real_name,omitempty"`       // 真实姓名
}

type UserInfoListQueryFilter struct {
	UserID      *int64         `query:"user_id" json:"user_id,omitempty"`           //
	Affiliation *string        `query:"affiliation" json:"affiliation,omitempty"`   // 工作单位
	Title       *string        `query:"title" json:"title,omitempty"`               // 职称
	IDCardType  *string        `query:"id_card_type" json:"id_card_type,omitempty"` // 证件类型
	IDCard      *string        `query:"id_card" json:"id_card,omitempty"`           //
	Biography   *string        `query:"biography" json:"biography,omitempty"`       // 自我介绍
	Tag         *string        `query:"tag" json:"tag,omitempty"`                   //
	Language    *string        `query:"language" json:"language,omitempty"`         //
	Gender      *consts.Gender `query:"gender" json:"gender,omitempty"`             // 性别
	Birthday    *time.Time     `query:"birthday" json:"birthday,omitempty"`         // 生日
	Education   *string        `query:"education" json:"education,omitempty"`       // 学历
	RealName    *string        `query:"real_name" json:"real_name,omitempty"`       // 真实姓名
}

type UserInfoItem struct {
	ID          int64         `json:"id,omitempty"`           //
	CreatedAt   time.Time     `json:"created_at,omitempty"`   //
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`   //
	UserID      int64         `json:"user_id,omitempty"`      //
	Affiliation string        `json:"affiliation,omitempty"`  // 工作单位
	Title       string        `json:"title,omitempty"`        // 职称
	IDCardType  string        `json:"id_card_type,omitempty"` // 证件类型
	IDCard      string        `json:"id_card,omitempty"`      //
	Biography   string        `json:"biography,omitempty"`    // 自我介绍
	Tag         string        `json:"tag,omitempty"`          //
	Language    string        `json:"language,omitempty"`     //
	Gender      consts.Gender `json:"gender,omitempty"`       // 性别
	Birthday    time.Time     `json:"birthday,omitempty"`     // 生日
	Education   string        `json:"education,omitempty"`    // 学历
	RealName    string        `json:"real_name,omitempty"`    // 真实姓名
}
