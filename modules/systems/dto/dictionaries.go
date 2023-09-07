package dto

import (
	"time"

	"github.com/atom-apps/door/common"
)

type DictionaryForm struct {
	Name        string            `form:"name" json:"name,omitempty"`               // 名称
	Slug        string            `form:"slug" json:"slug,omitempty"`               // 别名
	Description string            `form:"description" json:"description,omitempty"` // 描述
	Items       common.LabelItems `form:"items" json:"items,omitempty"`             // 选项
}

type DictionaryListQueryFilter struct {
	Name        *string `query:"name" json:"name,omitempty"`               // 名称
	Slug        *string `query:"slug" json:"slug,omitempty"`               // 别名
	Description *string `query:"description" json:"description,omitempty"` // 描述
}

type DictionaryItem struct {
	ID          uint64            `json:"id,omitempty"`          // ID
	CreatedAt   time.Time         `json:"created_at,omitempty"`  // 创建时间
	UpdatedAt   time.Time         `json:"updated_at,omitempty"`  // 更新时间
	Name        string            `json:"name,omitempty"`        // 名称
	Slug        string            `json:"slug,omitempty"`        // 别名
	Description string            `json:"description,omitempty"` // 描述
	Items       common.LabelItems `json:"items,omitempty"`       // 选项
}
