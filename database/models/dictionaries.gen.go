// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"github.com/atom-apps/door/common"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const TableNameDictionary = "dictionaries"

// Dictionary mapped from table <dictionaries>
type Dictionary struct {
	ID          uint64            `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`      // ID
	CreatedAt   time.Time         `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                      // 创建时间
	UpdatedAt   time.Time         `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`                      // 更新时间
	DeletedAt   gorm.DeletedAt    `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at" swaggertype:"string"` // 删除时间
	Name        string            `gorm:"column:name;type:varchar(198);not null;comment:名称" json:"name"`                          // 名称
	Slug        string            `gorm:"column:slug;type:varchar(120);not null;comment:别名" json:"slug"`                          // 别名
	Description string            `gorm:"column:description;type:varchar(198);not null;comment:描述" json:"description"`            // 描述
	Items       common.LabelItems `gorm:"column:items;type:varchar(198);not null;comment:选项" json:"items"`                        // 选项
}

func (*Dictionary) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameDictionary
	}
	return namer.TableName(TableNameDictionary)
}
