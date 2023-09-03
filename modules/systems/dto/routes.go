package dto

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
)

type RouteForm struct {
	Type     *consts.RouteType    `query:"type" json:"type,omitempty"`          //
	ParentID uint64               `form:"parent_id" json:"parent_id,omitempty"` //
	Name     string               `form:"name" json:"name,omitempty"`           //
	Path     string               `form:"path" json:"path,omitempty"`           //
	Metadata common.RouteMetadata `form:"metadata" json:"metadata,omitempty"`   //
}

type RouteListQueryFilter struct {
	Type     *consts.RouteType     `query:"type" json:"type,omitempty"`           //
	ParentID *uint64               `query:"parent_id" json:"parent_id,omitempty"` //
	Name     *string               `query:"name" json:"name,omitempty"`           //
	Path     *string               `query:"path" json:"path,omitempty"`           //
	Metadata *common.RouteMetadata `query:"metadata" json:"metadata,omitempty"`   //
}

type RouteItem struct {
	ID       uint64               `json:"id,omitempty"`                //
	Type     *consts.RouteType    `query:"type" json:"type,omitempty"` //
	ParentID uint64               `json:"parent_id,omitempty"`         //
	Name     string               `json:"name,omitempty"`              //
	Path     string               `json:"path,omitempty"`              //
	Metadata common.RouteMetadata `json:"metadata,omitempty"`          //
	Children []*RouteItem         `json:"children,omitempty"`          //
}
