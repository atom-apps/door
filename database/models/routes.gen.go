// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"github.com/atom-apps/door/common"
	"gorm.io/datatypes"
	"gorm.io/gorm/schema"
)

const TableNameRoute = "routes"

// Route mapped from table <routes>
type Route struct {
	ID       int64                                    `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Type     string                                   `gorm:"column:type;type:character varying(64);not null" json:"type"`
	ParentID int64                                    `gorm:"column:parent_id;type:bigint;not null" json:"parent_id"`
	Name     string                                   `gorm:"column:name;type:character varying(255);not null" json:"name"`
	Path     string                                   `gorm:"column:path;type:character varying(1024);not null" json:"path"`
	Metadata datatypes.JSONType[common.RouteMetadata] `gorm:"column:metadata;type:text" json:"metadata"`
}

func (*Route) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameRoute
	}
	return namer.TableName(TableNameRoute)
}
