// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm/schema"
)

const TableNameTenant = "tenants"

// Tenant mapped from table <tenants>
type Tenant struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID" json:"id"`       // ID
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp with time zone;comment:创建时间" json:"created_at"` // 创建时间
	Name        string    `gorm:"column:name;type:character varying(64);not null;comment:名称" json:"name"`         // 名称
	Description string    `gorm:"column:description;type:character varying(256);comment:描述" json:"description"`   // 描述
	Meta        string    `gorm:"column:meta;type:character varying(1024);comment:元数据" json:"meta"`               // 元数据
}

func (*Tenant) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameTenant
	}
	return namer.TableName(TableNameTenant)
}
