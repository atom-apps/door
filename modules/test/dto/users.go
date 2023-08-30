package dto

import (
	"time"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
)

type UserForm struct {
	UUID          string            `form:"uuid" json:"uuid,omitempty"`                     // UUID
	Username      string            `form:"username" json:"username,omitempty"`             // 用户名
	Password      string            `form:"password" json:"password,omitempty"`             // 密码
	Email         string            `form:"email" json:"email,omitempty"`                   // 邮箱
	EmailVerified bool              `form:"email_verified" json:"email_verified,omitempty"` // 邮箱是否验证
	Phone         string            `form:"phone" json:"phone,omitempty"`                   // 手机号
	DisplayName   string            `form:"display_name" json:"display_name,omitempty"`     // 显示名称
	Avatar        string            `form:"avatar" json:"avatar,omitempty"`                 // 头像
	Status        consts.UserStatus `form:"status" json:"status,omitempty"`                 // 状态
}

type UserListQueryFilter struct {
	UUID          *string            `query:"uuid" json:"uuid,omitempty"`                     // UUID
	Username      *string            `query:"username" json:"username,omitempty"`             // 用户名
	Password      *string            `query:"password" json:"password,omitempty"`             // 密码
	Email         *string            `query:"email" json:"email,omitempty"`                   // 邮箱
	EmailVerified *bool              `query:"email_verified" json:"email_verified,omitempty"` // 邮箱是否验证
	Phone         *string            `query:"phone" json:"phone,omitempty"`                   // 手机号
	DisplayName   *string            `query:"display_name" json:"display_name,omitempty"`     // 显示名称
	Avatar        *string            `query:"avatar" json:"avatar,omitempty"`                 // 头像
	Status        *consts.UserStatus `query:"status" json:"status,omitempty"`                 // 状态
}

func UserListQueryFilters() []common.Filter {
	return []common.Filter{
		{Type: common.FilterTypeString, Name: "id", Label: "ID"},
		{Type: common.FilterTypeString, Name: "created_at", Label: "创建时间"},
		{Type: common.FilterTypeString, Name: "updated_at", Label: "更新时间"},
		{Type: common.FilterTypeString, Name: "deleted_at", Label: "删除时间"},
		{Type: common.FilterTypeString, Name: "uuid", Label: "UUID"},
		{Type: common.FilterTypeString, Name: "username", Label: "用户名"},
		{Type: common.FilterTypeString, Name: "password", Label: "密码"},
		{Type: common.FilterTypeString, Name: "email", Label: "邮箱"},
		{Type: common.FilterTypeString, Name: "email_verified", Label: "邮箱是否验证"},
		{Type: common.FilterTypeString, Name: "phone", Label: "手机号"},
		{Type: common.FilterTypeString, Name: "display_name", Label: "显示名称"},
		{Type: common.FilterTypeString, Name: "avatar", Label: "头像"},
		{Type: common.FilterTypeString, Name: "status", Label: "状态"},
	}
}

type UserItem struct {
	ID            int64             `json:"id,omitempty"`             // ID
	CreatedAt     time.Time         `json:"created_at,omitempty"`     // 创建时间
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`     // 更新时间
	UUID          string            `json:"uuid,omitempty"`           // UUID
	Username      string            `json:"username,omitempty"`       // 用户名
	Password      string            `json:"password,omitempty"`       // 密码
	Email         string            `json:"email,omitempty"`          // 邮箱
	EmailVerified bool              `json:"email_verified,omitempty"` // 邮箱是否验证
	Phone         string            `json:"phone,omitempty"`          // 手机号
	DisplayName   string            `json:"display_name,omitempty"`   // 显示名称
	Avatar        string            `json:"avatar,omitempty"`         // 头像
	Status        consts.UserStatus `json:"status,omitempty"`         // 状态
}
