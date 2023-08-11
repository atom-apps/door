package dto

import (
	"time"
)

type SessionForm struct {
	UserID    int64  `form:"user_id" json:"user_id,omitempty"`       //
	SessionID string `form:"session_id" json:"session_id,omitempty"` //
}

type SessionListQueryFilter struct {
	UserID    *int64  `query:"user_id" json:"user_id,omitempty"`       //
	SessionID *string `query:"session_id" json:"session_id,omitempty"` //
}

type SessionItem struct {
	ID        int64     `json:"id,omitempty"`         //
	CreatedAt time.Time `json:"created_at,omitempty"` //
	UpdatedAt time.Time `json:"updated_at,omitempty"` //
	UserID    int64     `json:"user_id,omitempty"`    //
	SessionID string    `json:"session_id,omitempty"` //
}
