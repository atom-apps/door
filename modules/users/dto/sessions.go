package dto

import (
	"time"

	"github.com/atom-apps/door/database/models"
)

type SessionForm struct {
	UserID    uint64 `form:"user_id" json:"user_id,omitempty"`       //
	SessionID string `form:"session_id" json:"session_id,omitempty"` //
}

type SessionListQueryFilter struct {
	UserID    *uint64 `query:"user_id" json:"user_id,omitempty"`       //
	SessionID *string `query:"session_id" json:"session_id,omitempty"` //
}

type SessionItem struct {
	ID        uint64          `json:"id,omitempty"`         //
	CreatedAt time.Time       `json:"created_at,omitempty"` //
	UpdatedAt time.Time       `json:"updated_at,omitempty"` //
	ExpireAt  time.Time       `json:"expire_at,omitempty"`  //
	UserID    uint64          `json:"user_id,omitempty"`    //
	User      *models.User    `json:"user,omitempty"`       //
	SessionID string          `json:"session_id,omitempty"` //
	Tokens    []*models.Token `json:"tokens,omitempty"`     //
}
