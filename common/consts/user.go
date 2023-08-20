package consts

// swagger:enum UserStatus
// ENUM(default="", blocked)
type UserStatus string

const SessionName = "sessionid"

// swagger:enum Role
// ENUM(super_admin, tenant_admin, tenant_user)
type Role string
