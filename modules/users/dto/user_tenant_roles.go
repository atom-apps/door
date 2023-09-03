package dto

type UserTenantRoleForm struct {
	UserID   uint64 `form:"user_id" json:"user_id,omitempty"`     // 用户ID
	TenantID uint64 `form:"tenant_id" json:"tenant_id,omitempty"` // 租户ID
	RoleID   uint64 `form:"role_id" json:"role_id,omitempty"`     // 角色ID
}

type UserTenantRoleListQueryFilter struct {
	UserID   *uint64 `query:"user_id" json:"user_id,omitempty"`     // 用户ID
	TenantID *uint64 `query:"tenant_id" json:"tenant_id,omitempty"` // 租户ID
	RoleID   *uint64 `query:"role_id" json:"role_id,omitempty"`     // 角色ID
}

type UserTenantRoleItem struct {
	ID       uint64 `json:"id,omitempty"`        // ID
	UserID   uint64 `json:"user_id,omitempty"`   // 用户ID
	TenantID uint64 `json:"tenant_id,omitempty"` // 租户ID
	RoleID   uint64 `json:"role_id,omitempty"`   // 角色ID
}
