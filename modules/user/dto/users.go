package dto

import (
	"time"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
)

type UserForm struct {
	UUID          string            `form:"uuid" json:"uuid,omitempty"`                     //
	Username      string            `form:"username" json:"username,omitempty"`             //
	Password      string            `form:"password" json:"password,omitempty"`             //
	Email         string            `form:"email" json:"email,omitempty"`                   //
	EmailVerified bool              `form:"email_verified" json:"email_verified,omitempty"` //
	Phone         string            `form:"phone" json:"phone,omitempty"`                   //
	DisplayName   string            `form:"display_name" json:"display_name,omitempty"`     //
	Avatar        string            `form:"avatar" json:"avatar,omitempty"`                 //
	Status        consts.UserStatus `json:"status,omitempty"`                               //
}

type UserListQueryFilter struct {
	IDs           []int64            `query:"ids" json:"ids,omitempty"`                       //
	UUID          *string            `query:"uuid" json:"uuid,omitempty"`                     //
	Username      *string            `query:"username" json:"username,omitempty"`             //
	Email         *string            `query:"email" json:"email,omitempty"`                   //
	EmailVerified *bool              `query:"email_verified" json:"email_verified,omitempty"` //
	Phone         *string            `query:"phone" json:"phone,omitempty"`                   //
	DisplayName   *string            `query:"display_name" json:"display_name,omitempty"`     //
	Status        *consts.UserStatus `json:"status,omitempty"`                                //
}

func UserListQueryFilters() []common.Filter {
	return []common.Filter{
		{Type: common.FilterTypeList, Name: "ids", Label: "ID"},
		{Type: common.FilterTypeString, Name: "uuid", Label: "UUID"},
		{Type: common.FilterTypeString, Name: "username", Label: "用户名"},
		{Type: common.FilterTypeString, Name: "display_name", Label: "昵称"},
		{Type: common.FilterTypeString, Name: "email", Label: "邮箱"},
		{Type: common.FilterTypeBool, Name: "email_verified", Label: "已验证"},
		{Type: common.FilterTypeString, Name: "phone", Label: "手机号"},
		{Type: common.FilterTypeString, Name: "status", Value: string(consts.UserStatusDefault), Items: consts.UserStatusLabel(false), Label: "状态"},
	}
}

type UserItem struct {
	ID            int64             `json:"id,omitempty"`             //
	CreatedAt     time.Time         `json:"created_at,omitempty"`     //
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`     //
	UUID          string            `json:"uuid,omitempty"`           //
	Username      string            `json:"username,omitempty"`       //
	Email         string            `json:"email,omitempty"`          //
	EmailVerified bool              `json:"email_verified,omitempty"` //
	Phone         string            `json:"phone,omitempty"`          //
	DisplayName   string            `json:"display_name,omitempty"`   //
	Avatar        string            `json:"avatar,omitempty"`         //
	Status        consts.UserStatus `json:"status,omitempty"`         //
}
