package dto

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
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

type UserItem struct {
	ID            int64                 `json:"id,omitempty"`             //
	CreatedAt     time.Time             `json:"created_at,omitempty"`     //
	UpdatedAt     time.Time             `json:"updated_at,omitempty"`     //
	UUID          string                `json:"uuid,omitempty"`           //
	Username      string                `json:"username,omitempty"`       //
	Email         string                `json:"email,omitempty"`          //
	EmailVerified bool                  `json:"email_verified,omitempty"` //
	Phone         string                `json:"phone,omitempty"`          //
	DisplayName   string                `json:"display_name,omitempty"`   //
	Avatar        string                `json:"avatar,omitempty"`         //
	Status        consts.UserStatus     `json:"status,omitempty"`         //
	TenantRoles   []*UserItemTenantRole `json:"tenant_roles,omitempty"`
}

type UserItemTenantRole struct {
	Role   *models.Role   `json:"role,omitempty"`
	Tenant *models.Tenant `json:"tenant,omitempty"`
}
