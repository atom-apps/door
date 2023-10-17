package dto

import (
	"time"
)

type BookForm struct {
	TenantID        uint64 `form:"tenant_id" json:"tenant_id,omitempty"`                 // 租户ID
	UserID          uint64 `form:"user_id" json:"user_id,omitempty"`                     // 用户ID
	Title           string `form:"title" json:"title,omitempty"`                         // 书名
	ThumbnailFileID uint64 `form:"thumbnail_file_id" json:"thumbnail_file_id,omitempty"` // 书名
	Description     string `form:"description" json:"description,omitempty"`             // 简介
	Content         string `form:"content" json:"content,omitempty"`                     // 详细介绍
	Author          string `form:"author" json:"author,omitempty"`                       // 原作者
	Source          string `form:"source" json:"source,omitempty"`                       // 原书地址
	Isbn            string `form:"isbn" json:"isbn,omitempty"`                           // ISBN
	Price           uint64 `form:"price" json:"price,omitempty"`                         // 价格
}

type BookListQueryFilter struct {
	TenantID    *uint64 `query:"tenant_id" json:"tenant_id,omitempty"`     // 租户ID
	UserID      *uint64 `query:"user_id" json:"user_id,omitempty"`         // 用户ID
	Title       *string `query:"title" json:"title,omitempty"`             // 书名
	Description *string `query:"description" json:"description,omitempty"` // 简介
	Content     *string `query:"content" json:"content,omitempty"`         // 详细介绍
	Author      *string `query:"author" json:"author,omitempty"`           // 原作者
	Source      *string `query:"source" json:"source,omitempty"`           // 原书地址
	Isbn        *string `query:"isbn" json:"isbn,omitempty"`               // ISBN
	Price       *uint64 `query:"price" json:"price,omitempty"`             // 价格
}

type BookItem struct {
	ID              uint64    `json:"id,omitempty"`                                         // ID
	CreatedAt       time.Time `json:"created_at,omitempty"`                                 // 创建时间
	UpdatedAt       time.Time `json:"updated_at,omitempty"`                                 // 更新时间
	TenantID        uint64    `json:"tenant_id,omitempty"`                                  // 租户ID
	UserID          uint64    `json:"user_id,omitempty"`                                    // 用户ID
	Title           string    `json:"title,omitempty"`                                      // 书名
	ThumbnailFileID uint64    `form:"thumbnail_file_id" json:"thumbnail_file_id,omitempty"` // 书名
	Thumbnail       string    `form:"thumbnail" json:"thumbnail,omitempty"`                 // 封面
	Description     string    `json:"description,omitempty"`                                // 简介
	Content         string    `json:"content,omitempty"`                                    // 详细介绍
	Author          string    `json:"author,omitempty"`                                     // 原作者
	Source          string    `json:"source,omitempty"`                                     // 原书地址
	Isbn            string    `json:"isbn,omitempty"`                                       // ISBN
	Price           uint64    `json:"price,omitempty"`                                      // 价格
	ChapterCount    int64     `json:"chapter_count"`                                        // 章节数量
	ArticleCount    int64     `json:"article_count"`                                        // 文章数量
}
