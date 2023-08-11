package dto

import (
	"time"
)

type TenantUserForm struct {
	TenantID int64 `form:"tenant_id" json:"tenant_id,omitempty"` //
	UserID   int64 `form:"user_id" json:"user_id,omitempty"`     //
	IsAdmin  bool  `form:"is_admin" json:"is_admin,omitempty"`   //
}

type TenantUserListQueryFilter struct {
	TenantID *int64 `query:"tenant_id" json:"tenant_id,omitempty"` //
	UserID   *int64 `query:"user_id" json:"user_id,omitempty"`     //
	IsAdmin  *bool  `query:"is_admin" json:"is_admin,omitempty"`   //
}

type TenantUserItem struct {
	ID        int64     `json:"id,omitempty"`         //
	CreatedAt time.Time `json:"created_at,omitempty"` //
	UpdatedAt time.Time `json:"updated_at,omitempty"` //
	TenantID  int64     `json:"tenant_id,omitempty"`  //
	UserID    int64     `json:"user_id,omitempty"`    //
	IsAdmin   bool      `json:"is_admin,omitempty"`   //
}
