// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm/schema"
)

const TableNameArticleForwardSource = "article_forward_sources"

// ArticleForwardSource mapped from table <article_forward_sources>
type ArticleForwardSource struct {
	ID           uint64    `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                     // 创建时间
	ArticleID    uint64    `gorm:"column:article_id;type:bigint(20) unsigned;comment:文章ID" json:"article_id"`             // 文章ID
	Source       string    `gorm:"column:source;type:varchar(1024);comment:原文地址" json:"source"`                           // 原文地址
	SourceAuthor string    `gorm:"column:source_author;type:varchar(64);comment:原文作者" json:"source_author"`               // 原文作者
}

func (*ArticleForwardSource) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameArticleForwardSource
	}
	return namer.TableName(TableNameArticleForwardSource)
}
