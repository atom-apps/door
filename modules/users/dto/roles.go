package dto

type RoleForm struct {
	Name        string `form:"name" json:"name,omitempty"`               //
	Slug        string `form:"slug" json:"slug,omitempty"`               //
	Description string `form:"description" json:"description,omitempty"` //
	ParentID    int64  `form:"parent_id" json:"parent_id,omitempty"`     //
}

type RoleListQueryFilter struct {
	Name        *string `query:"name" json:"name,omitempty"`               //
	Slug        *string `form:"slug" json:"slug,omitempty"`                //
	Description *string `query:"description" json:"description,omitempty"` //
	ParentID    *int64  `query:"parent_id" json:"parent_id,omitempty"`     //
}

type RoleItem struct {
	ID          int64     `json:"id,omitempty"`               //
	Name        string    `json:"name,omitempty"`             //
	Slug        string    `form:"slug" json:"slug,omitempty"` //
	Description string    `json:"description,omitempty"`      //
	ParentID    int64     `json:"parent_id,omitempty"`        //
	Parent      *RoleItem `json:"parent,omitempty"`           //
}
