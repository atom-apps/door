package dto

type PermissionRuleForm struct {
	Ptype string `form:"ptype" json:"ptype,omitempty"` //
	V0    string `form:"v0" json:"v0,omitempty"`       //
	V1    string `form:"v1" json:"v1,omitempty"`       //
	V2    string `form:"v2" json:"v2,omitempty"`       //
	V3    string `form:"v3" json:"v3,omitempty"`       //
	V4    string `form:"v4" json:"v4,omitempty"`       //
	V5    string `form:"v5" json:"v5,omitempty"`       //
}

type PermissionRuleListQueryFilter struct {
	Ptype *string `query:"ptype" json:"ptype,omitempty"` //
	V0    *string `query:"v0" json:"v0,omitempty"`       //
	V1    *string `query:"v1" json:"v1,omitempty"`       //
	V2    *string `query:"v2" json:"v2,omitempty"`       //
	V3    *string `query:"v3" json:"v3,omitempty"`       //
	V4    *string `query:"v4" json:"v4,omitempty"`       //
	V5    *string `query:"v5" json:"v5,omitempty"`       //
}

type PermissionRuleItem struct {
	ID    int64  `json:"id,omitempty"`    //
	Ptype string `json:"ptype,omitempty"` //
	V0    string `json:"v0,omitempty"`    //
	V1    string `json:"v1,omitempty"`    //
	V2    string `json:"v2,omitempty"`    //
	V3    string `json:"v3,omitempty"`    //
	V4    string `json:"v4,omitempty"`    //
	V5    string `json:"v5,omitempty"`    //
}
