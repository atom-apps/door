package dto

type LocationForm struct {
	Code     uint64 `form:"code" json:"code,omitempty"`         // 行政区划代码
	Name     string `form:"name" json:"name,omitempty"`         //
	Province string `form:"province" json:"province,omitempty"` //
	City     string `form:"city" json:"city,omitempty"`         //
	Area     string `form:"area" json:"area,omitempty"`         //
	Town     string `form:"town" json:"town,omitempty"`         //
}

type LocationListQueryFilter struct {
	Code     *uint64 `query:"code" json:"code,omitempty"`         // 行政区划代码
	Name     *string `query:"name" json:"name,omitempty"`         //
	Province *string `query:"province" json:"province,omitempty"` //
	City     *string `query:"city" json:"city,omitempty"`         //
	Area     *string `query:"area" json:"area,omitempty"`         //
	Town     *string `query:"town" json:"town,omitempty"`         //
}

type LocationItem struct {
	ID       uint64 `json:"id,omitempty"`       //
	Code     int64  `json:"code,omitempty"`     // 行政区划代码
	Name     string `json:"name,omitempty"`     //
	Province string `json:"province,omitempty"` //
	City     string `json:"city,omitempty"`     //
	Area     string `json:"area,omitempty"`     //
	Town     string `json:"town,omitempty"`     //
}
