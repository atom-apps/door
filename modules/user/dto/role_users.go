package dto

import (
	"time"
)

type RoleUserForm struct {
	RoleID   int64 `form:"role_id" json:"role_id,omitempty"`     //
	UserID   int64 `form:"user_id" json:"user_id,omitempty"`     //
	TenantID int64 `form:"tenant_id" json:"tenant_id,omitempty"` //
}

type RoleUserListQueryFilter struct {
	RoleID   *int64 `query:"role_id" json:"role_id,omitempty"`     //
	UserID   *int64 `query:"user_id" json:"user_id,omitempty"`     //
	TenantID *int64 `query:"tenant_id" json:"tenant_id,omitempty"` //
}

type RoleUserItem struct {
	ID        int64     `json:"id,omitempty"`         //
	CreatedAt time.Time `json:"created_at,omitempty"` //
	UpdatedAt time.Time `json:"updated_at,omitempty"` //
	RoleID    int64     `json:"role_id,omitempty"`    //
	UserID    int64     `json:"user_id,omitempty"`    //
	TenantID  int64     `json:"tenant_id,omitempty"`  //
}
