package dto

import "github.com/atom-apps/door/database/models"

type RoleForm struct {
	Name        string `form:"name" json:"name,omitempty"`               //
	Slug        string `form:"slug" json:"slug,omitempty"`               //
	Description string `form:"description" json:"description,omitempty"` //
	ParentID    uint64 `form:"parent_id" json:"parent_id,omitempty"`     //
}

type RoleListQueryFilter struct {
	Name        *string `query:"name" json:"name,omitempty"`               //
	Slug        *string `form:"slug" json:"slug,omitempty"`                //
	Description *string `query:"description" json:"description,omitempty"` //
	ParentID    *uint64 `query:"parent_id" json:"parent_id,omitempty"`     //
}

type RoleItem struct {
	ID          uint64              `json:"id,omitempty"`               //
	Name        string              `json:"name,omitempty"`             //
	Slug        string              `form:"slug" json:"slug,omitempty"` //
	Description string              `json:"description,omitempty"`      //
	ParentID    uint64              `json:"parent_id,omitempty"`        //
	Parent      *RoleItem           `json:"parent,omitempty"`           //
	UserAmount  int64               `json:"user_amount"`                //
	Tenants     []*models.Tenant    `json:"tenants,omitempty"`          //
	Permissions map[uint64][]uint64 `json:"permissions,omitempty"`      //
}
