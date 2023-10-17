package dto

import (
	"time"
)

type ChapterForm struct {
	TenantID    uint64 `form:"tenant_id" json:"tenant_id,omitempty"`     // 租户ID
	UserID      uint64 `form:"user_id" json:"user_id,omitempty"`         // 用户ID
	BookID      uint64 `form:"book_id" json:"book_id,omitempty"`         // 书ID
	Title       string `form:"title" json:"title,omitempty"`             // 章节名
	Description string `form:"description" json:"description,omitempty"` // 简介
	Content     string `form:"content" json:"content,omitempty"`         // 详细介绍
}

type ChapterListQueryFilter struct {
	TenantID    *uint64 `query:"tenant_id" json:"tenant_id,omitempty"`     // 租户ID
	UserID      *uint64 `query:"user_id" json:"user_id,omitempty"`         // 用户ID
	BookID      *uint64 `query:"book_id" json:"book_id,omitempty"`         // 书ID
	Title       *string `query:"title" json:"title,omitempty"`             // 章节名
	Description *string `query:"description" json:"description,omitempty"` // 简介
	Content     *string `query:"content" json:"content,omitempty"`         // 详细介绍
}

type ChapterItem struct {
	ID           uint64    `json:"id,omitempty"`          // ID
	CreatedAt    time.Time `json:"created_at,omitempty"`  // 创建时间
	UpdatedAt    time.Time `json:"updated_at,omitempty"`  // 更新时间
	TenantID     uint64    `json:"tenant_id,omitempty"`   // 租户ID
	UserID       uint64    `json:"user_id,omitempty"`     // 用户ID
	BookID       uint64    `json:"book_id,omitempty"`     // 书ID
	Title        string    `json:"title,omitempty"`       // 章节名
	Description  string    `json:"description,omitempty"` // 简介
	Content      string    `json:"content,omitempty"`     // 详细介绍
	ArticleCount int64     `json:"article_count"`         // 文章数量
}
