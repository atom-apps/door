package dto

import "github.com/atom-apps/door/common/model"

type RouteForm struct {
	ParentID uint64              `form:"parent_id" json:"parent_id,omitempty"` //
	Name     string              `form:"name" json:"name,omitempty"`           //
	Path     string              `form:"path" json:"path,omitempty"`           //
	Metadata model.RouteMetadata `form:"metadata" json:"metadata,omitempty"`   //
}

type RouteListQueryFilter struct {
	ParentID *uint64              `query:"parent_id" json:"parent_id,omitempty"` //
	Name     *string              `query:"name" json:"name,omitempty"`           //
	Path     *string              `query:"path" json:"path,omitempty"`           //
	Metadata *model.RouteMetadata `query:"metadata" json:"metadata,omitempty"`   //
}

type RouteItem struct {
	ID       uint64              `json:"id,omitempty"`        //
	ParentID uint64              `json:"parent_id,omitempty"` //
	Name     string              `json:"name,omitempty"`      //
	Path     string              `json:"path,omitempty"`      //
	Metadata model.RouteMetadata `json:"metadata,omitempty"`  //
	Children []*RouteItem        `json:"children,omitempty"`  //
}
