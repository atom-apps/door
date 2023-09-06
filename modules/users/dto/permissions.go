package dto

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
)

type PermissionForm struct {
	TenantID uint64 `form:"tenant_id" json:"tenant_id,omitempty"` // 租户ID
	RoleID   uint64 `form:"role_id" json:"role_id,omitempty"`     // 角色ID
	Path     string `form:"path" json:"path,omitempty"`           // 路由
	Action   string `form:"action" json:"action,omitempty"`       // 请求方式
}

type PermissionListQueryFilter struct {
	TenantID *uint64 `query:"tenant_id" json:"tenant_id,omitempty"` // 租户ID
	RoleID   *uint64 `query:"role_id" json:"role_id,omitempty"`     // 角色ID
	Path     *string `query:"path" json:"path,omitempty"`           // 路由
	Action   *string `query:"action" json:"action,omitempty"`       // 请求方式
}

type PermissionItem struct {
	ID       uint64         `json:"id,omitempty"`        // ID
	TenantID uint64         `json:"tenant_id,omitempty"` // 租户ID
	Tenant   *models.Tenant `json:"tenant,omitempty"`    // 租户
	RoleID   uint64         `json:"role_id,omitempty"`   // 角色ID
	Role     *models.Role   `json:"role,omitempty"`      // 角色
	Path     string         `json:"path,omitempty"`      // 路由
	Action   string         `json:"action,omitempty"`    // 请求方式
}

type PermissionTree struct {
	ID       uint64               `json:"id"`
	Name     string               `json:"name"`
	Method   string               `json:"method"`
	Path     string               `json:"path"`
	ParentID uint64               `json:"parent_id"`
	Metadata common.RouteMetadata `json:"metadata"`
	Children []*PermissionTree    `json:"children,omitempty"`
}
