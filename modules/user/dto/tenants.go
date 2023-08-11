package dto

import (
	"time"
)

type TenantForm struct {
	Name        string `form:"name" json:"name,omitempty"`               //
	Description string `form:"description" json:"description,omitempty"` //
	Meta        string `form:"meta" json:"meta,omitempty"`               //
}

type TenantListQueryFilter struct {
	Name        *string `query:"name" json:"name,omitempty"`               //
	Description *string `query:"description" json:"description,omitempty"` //
	Meta        *string `query:"meta" json:"meta,omitempty"`               //
}

type TenantItem struct {
	ID          int64     `json:"id,omitempty"`          //
	CreatedAt   time.Time `json:"created_at,omitempty"`  //
	UpdatedAt   time.Time `json:"updated_at,omitempty"`  //
	Name        string    `json:"name,omitempty"`        //
	Description string    `json:"description,omitempty"` //
	Meta        string    `json:"meta,omitempty"`        //
}
