// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"gorm.io/datatypes"
	"gorm.io/gorm/schema"
)

const TableNameRoute = "routes"

// Route mapped from table <routes>
type Route struct {
	ID       uint64                                   `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Type     consts.RouteType                         `gorm:"column:type;type:varchar(64);not null" json:"type"`
	ParentID uint64                                   `gorm:"column:parent_id;type:bigint unsigned;not null" json:"parent_id"`
	Name     string                                   `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Path     string                                   `gorm:"column:path;type:varchar(1024);not null" json:"path"`
	Metadata datatypes.JSONType[common.RouteMetadata] `gorm:"column:metadata;type:varchar(191)" json:"metadata"`
	Order    uint64                                   `gorm:"column:order;type:bigint unsigned" json:"order"`
}

func (*Route) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameRoute
	}
	return namer.TableName(TableNameRoute)
}
