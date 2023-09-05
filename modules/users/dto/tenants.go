package dto

import (
	"time"

	"github.com/atom-apps/door/database/models"
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
	ID          uint64         `json:"id,omitempty"`          //
	CreatedAt   time.Time      `json:"created_at,omitempty"`  //
	Name        string         `json:"name,omitempty"`        //
	Description string         `json:"description,omitempty"` //
	Meta        string         `json:"meta,omitempty"`        //
	UserAmount  int64          `json:"user_amount,omitempty"` //
	Roles       []*models.Role `json:"roles,omitempty"`       //
}
