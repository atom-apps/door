// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"gorm.io/gorm/schema"
)

const TableNamePermissionRule = "permission_rules"

// PermissionRule mapped from table <permission_rules>
type PermissionRule struct {
	ID    uint64 `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Ptype string `gorm:"column:ptype;type:varchar(191)" json:"ptype"`
	V0    string `gorm:"column:v0;type:varchar(191)" json:"v0"`
	V1    string `gorm:"column:v1;type:varchar(191)" json:"v1"`
	V2    string `gorm:"column:v2;type:varchar(191)" json:"v2"`
	V3    string `gorm:"column:v3;type:varchar(191)" json:"v3"`
	V4    string `gorm:"column:v4;type:varchar(191)" json:"v4"`
	V5    string `gorm:"column:v5;type:varchar(191)" json:"v5"`
}

func (*PermissionRule) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNamePermissionRule
	}
	return namer.TableName(TableNamePermissionRule)
}
