package dto

import (
	"time"

	"github.com/atom-apps/door/common"
)

type RoleForm struct {
	Name        string `form:"name" json:"name,omitempty"`               // 名称
	Slug        string `form:"slug" json:"slug,omitempty"`               // 唯一标识
	Description string `form:"description" json:"description,omitempty"` // 描述
	ParentID    int64  `form:"parent_id" json:"parent_id,omitempty"`     // 父角色
}

type RoleListQueryFilter struct {
	Name        *string `query:"name" json:"name,omitempty"`               // 名称
	Slug        *string `query:"slug" json:"slug,omitempty"`               // 唯一标识
	Description *string `query:"description" json:"description,omitempty"` // 描述
	ParentID    *int64  `query:"parent_id" json:"parent_id,omitempty"`     // 父角色
}

func RoleListQueryFilters() []common.Filter {
	return []common.Filter{
		{Type: common.FilterTypeString, Name: "id", Label: "ID"},
		{Type: common.FilterTypeString, Name: "created_at", Label: "创建时间"},
		{Type: common.FilterTypeString, Name: "name", Label: "名称"},
		{Type: common.FilterTypeString, Name: "slug", Label: "唯一标识"},
		{Type: common.FilterTypeString, Name: "description", Label: "描述"},
		{Type: common.FilterTypeString, Name: "parent_id", Label: "父角色"},
	}
}

type RoleItem struct {
	ID          int64     `json:"id,omitempty"`          // ID
	CreatedAt   time.Time `json:"created_at,omitempty"`  // 创建时间
	Name        string    `json:"name,omitempty"`        // 名称
	Slug        string    `json:"slug,omitempty"`        // 唯一标识
	Description string    `json:"description,omitempty"` // 描述
	ParentID    int64     `json:"parent_id,omitempty"`   // 父角色
}
