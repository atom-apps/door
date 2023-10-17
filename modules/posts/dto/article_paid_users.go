package dto

import (
	"time"
)

type ArticlePaidUserForm struct {
	TenantID  uint64 `form:"tenant_id" json:"tenant_id,omitempty"`   // 租户ID
	UserID    uint64 `form:"user_id" json:"user_id,omitempty"`       // 用户ID
	ArticleID uint64 `form:"article_id" json:"article_id,omitempty"` // 文章ID
	PriceType int64  `form:"price_type" json:"price_type,omitempty"` // 付费类型
	Price     uint64 `form:"price" json:"price,omitempty"`           // 付费价格
	EndAt     uint64 `form:"end_at" json:"end_at,omitempty"`         // 付费结束时间
}

type ArticlePaidUserListQueryFilter struct {
	TenantID  *uint64 `query:"tenant_id" json:"tenant_id,omitempty"`   // 租户ID
	UserID    *uint64 `query:"user_id" json:"user_id,omitempty"`       // 用户ID
	ArticleID *uint64 `query:"article_id" json:"article_id,omitempty"` // 文章ID
	PriceType *int64  `query:"price_type" json:"price_type,omitempty"` // 付费类型
	Price     *uint64 `query:"price" json:"price,omitempty"`           // 付费价格
	EndAt     *uint64 `query:"end_at" json:"end_at,omitempty"`         // 付费结束时间
}

type ArticlePaidUserItem struct {
	ID        uint64    `json:"id,omitempty"`         // ID
	CreatedAt time.Time `json:"created_at,omitempty"` // 创建时间
	TenantID  uint64    `json:"tenant_id,omitempty"`  // 租户ID
	UserID    uint64    `json:"user_id,omitempty"`    // 用户ID
	ArticleID uint64    `json:"article_id,omitempty"` // 文章ID
	PriceType int64     `json:"price_type,omitempty"` // 付费类型
	Price     uint64    `json:"price,omitempty"`      // 付费价格
	EndAt     uint64    `json:"end_at,omitempty"`     // 付费结束时间
}
