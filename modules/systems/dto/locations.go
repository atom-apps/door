package dto

import (
	"time"
)

type LocationForm struct {
	Code     string `form:"code" json:"code,omitempty"`         // 行政区划代码
	Name     string `form:"name" json:"name,omitempty"`         // 名称
	Province string `form:"province" json:"province,omitempty"` // 省/直辖市
	City     string `form:"city" json:"city,omitempty"`         // 市
	Area     string `form:"area" json:"area,omitempty"`         // 区县
	Town     string `form:"town" json:"town,omitempty"`         // 乡镇
}

type LocationListQueryFilter struct {
	Code     *string `query:"code" json:"code,omitempty"`         // 行政区划代码
	Name     *string `query:"name" json:"name,omitempty"`         // 名称
	Province *string `query:"province" json:"province,omitempty"` // 省/直辖市
	City     *string `query:"city" json:"city,omitempty"`         // 市
	Area     *string `query:"area" json:"area,omitempty"`         // 区县
	Town     *string `query:"town" json:"town,omitempty"`         // 乡镇
}

type LocationItem struct {
	ID        uint64    `json:"id,omitempty"`         // ID
	CreatedAt time.Time `json:"created_at,omitempty"` // 创建时间
	Code      string    `json:"code,omitempty"`       // 行政区划代码
	Name      string    `json:"name,omitempty"`       // 名称
	Province  string    `json:"province,omitempty"`   // 省/直辖市
	City      string    `json:"city,omitempty"`       // 市
	Area      string    `json:"area,omitempty"`       // 区县
	Town      string    `json:"town,omitempty"`       // 乡镇
}

type LocationDetail struct {
	Code     string `json:"code,omitempty"`      // 行政区划代码
	TownCode string `json:"town_code,omitempty"` //
	Province string `json:"province,omitempty"`  // 省/直辖市
	City     string `json:"city,omitempty"`      // 市
	Area     string `json:"area,omitempty"`      // 区县
	Town     string `json:"town,omitempty"`      // 乡镇
}
