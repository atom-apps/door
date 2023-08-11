// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const TableNameUserInfo = "user_infos"

// UserInfo mapped from table <user_infos>
type UserInfo struct {
	ID          int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp with time zone" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp with time zone" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone" json:"deleted_at"`
	UserID      int64          `gorm:"column:user_id;type:bigint" json:"user_id"`
	Affiliation string         `gorm:"column:affiliation;type:character varying(128);comment:工作单位" json:"affiliation"`            // 工作单位
	Title       string         `gorm:"column:title;type:character varying(128);comment:职称" json:"title"`                          // 职称
	IDCardType  string         `gorm:"column:id_card_type;type:character varying(128);not null;comment:证件类型" json:"id_card_type"` // 证件类型
	IDCard      string         `gorm:"column:id_card;type:character varying(128)" json:"id_card"`
	Biography   string         `gorm:"column:biography;type:character varying(256);comment:自我介绍" json:"biography"` // 自我介绍
	Tag         string         `gorm:"column:tag;type:character varying(128)" json:"tag"`
	Language    string         `gorm:"column:language;type:character varying(128)" json:"language"`
	Gender      consts.Gender  `gorm:"column:gender;type:text;not null;comment:性别" json:"gender"`                  // 性别
	Birthday    time.Time      `gorm:"column:birthday;type:timestamp with time zone;comment:生日" json:"birthday"`   // 生日
	Education   string         `gorm:"column:education;type:character varying(128);comment:学历" json:"education"`   // 学历
	RealName    string         `gorm:"column:real_name;type:character varying(128);comment:真实姓名" json:"real_name"` // 真实姓名
}

func (*UserInfo) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameUserInfo
	}
	return namer.TableName(TableNameUserInfo)
}
