package dto

type PermissionCheckForm struct {
	Path   string `json:"path,omitempty"`
	Method string `json:"method,omitempty"`
}
