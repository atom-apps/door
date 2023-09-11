package dto

import (
	"time"
)

type UserAddressForm struct {
	Code    string `form:"code" json:"code,omitempty"`         // 行政区划代码
	Town    string `form:"town" json:"town,omitempty"`         // 街道
	Detail  string `form:"detail" json:"detail,omitempty"`     // 详细地址
	Name    string `form:"name" json:"name,omitempty"`         // 姓名
	Phone   string `form:"phone" json:"phone,omitempty"`       // 联系电话
	ZipCode string `form:"zip_code" json:"zip_code,omitempty"` // 邮编
}

type UserAddressListQueryFilter struct {
	UserID  *uint64 `query:"user_id" json:"user_id,omitempty"`   // 用户ID
	Code    *string `query:"code" json:"code,omitempty"`         // 行政区划代码
	Town    *string `query:"town" json:"town,omitempty"`         // 街道
	Detail  *string `query:"detail" json:"detail,omitempty"`     // 详细地址
	Name    *string `query:"name" json:"name,omitempty"`         // 姓名
	Phone   *string `query:"phone" json:"phone,omitempty"`       // 联系电话
	ZipCode *string `query:"zip_code" json:"zip_code,omitempty"` // 邮编
}

type UserAddressItem struct {
	ID        uint64    `json:"id,omitempty"`         // ID
	CreatedAt time.Time `json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `json:"updated_at,omitempty"` // 更新时间
	UserID    uint64    `json:"user_id,omitempty"`    // 用户ID
	Code      string    `json:"code,omitempty"`       // 行政区划代码
	Town      string    `json:"town,omitempty"`       // 街道
	Detail    string    `json:"detail,omitempty"`     // 详细地址
	Name      string    `json:"name,omitempty"`       // 姓名
	Phone     string    `json:"phone,omitempty"`      // 联系电话
	ZipCode   string    `json:"zip_code,omitempty"`   // 邮编
}
