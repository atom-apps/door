// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm/schema"
)

const TableNameLocation = "locations"

// Location mapped from table <locations>
type Location struct {
	ID        uint64    `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                     // 创建时间
	Code      string    `gorm:"column:code;type:varchar(6);comment:行政区划代码" json:"code"`                                // 行政区划代码
	Name      string    `gorm:"column:name;type:varchar(128);comment:名称" json:"name"`                                  // 名称
	Province  string    `gorm:"column:province;type:varchar(2);comment:省/直辖市" json:"province"`                         // 省/直辖市
	City      string    `gorm:"column:city;type:varchar(2);comment:市" json:"city"`                                     // 市
	Area      string    `gorm:"column:area;type:varchar(2);comment:区县" json:"area"`                                    // 区县
	Town      string    `gorm:"column:town;type:varchar(12);comment:乡镇" json:"town"`                                   // 乡镇
}

func (*Location) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameLocation
	}
	return namer.TableName(TableNameLocation)
}
