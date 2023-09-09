package dto

import (
	"time"

	"github.com/atom-apps/door/common"
)

type MenuForm struct {
	Name     string              `form:"name" json:"name,omitempty"`           // 名称
	Slug     string              `form:"slug" json:"slug,omitempty"`           // 别名
	ParentID uint64              `form:"parent_id" json:"parent_id,omitempty"` // 父ID
	Metadata common.MenuMetadata `form:"metadata" json:"metadata,omitempty"`   // 元数据
}

type MenuListQueryFilter struct {
	Name     *string `query:"name" json:"name,omitempty"`           // 名称
	Slug     *string `query:"slug" json:"slug,omitempty"`           // 别名
	GroupID  *uint64 `query:"group_id" json:"group_id,omitempty"`   // 组
	ParentID *uint64 `query:"parent_id" json:"parent_id,omitempty"` // 父ID
	Metadata *string `query:"metadata" json:"metadata,omitempty"`   // 元数据
}

type MenuItem struct {
	ID        uint64              `json:"id,omitempty"`         // ID
	CreatedAt time.Time           `json:"created_at,omitempty"` // 创建时间
	Name      string              `json:"name,omitempty"`       // 名称
	Slug      string              `json:"slug,omitempty"`       // 别名
	GroupID   uint64              `json:"group_id,omitempty"`   // 组
	ParentID  uint64              `json:"parent_id,omitempty"`  // 父ID
	Metadata  common.MenuMetadata `json:"metadata,omitempty"`   // 元数据
	Children  []MenuItem          `json:"children,omitempty"`   //
}

type MenuTreeItem struct {
	Key      string          `json:"key,omitempty"`      // ID
	Title    string          `json:"title,omitempty"`    // 名称
	Children []*MenuTreeItem `json:"children,omitempty"` //
}
