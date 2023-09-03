package dto

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
)

type TokenForm struct {
	UserID        uint64    `form:"user_id" json:"user_id,omitempty"`                 //
	TokenExpireAt time.Time `form:"token_expire_at" json:"token_expire_at,omitempty"` //
	AccessToken   string    `form:"access_token" json:"access_token,omitempty"`       //
	RefreshToken  string    `form:"refresh_token" json:"refresh_token,omitempty"`     //
	Scope         string    `form:"scope" json:"scope,omitempty"`                     //
	TokenType     string    `form:"token_type" json:"token_type,omitempty"`           //
	CodeChallenge string    `form:"code_challenge" json:"code_challenge,omitempty"`   //
	Code          string    `form:"code" json:"code,omitempty"`                       //
	CodeExpireAt  time.Time `form:"code_expire_at" json:"code_expire_at,omitempty"`   //
	Used          bool      `form:"used" json:"used,omitempty"`                       //
}

type TokenListQueryFilter struct {
	UserID        *uint64           `query:"user_id" json:"user_id,omitempty"`                 //
	TokenExpireAt *time.Time        `query:"token_expire_at" json:"token_expire_at,omitempty"` //
	AccessToken   *string           `query:"access_token" json:"access_token,omitempty"`       //
	RefreshToken  *string           `query:"refresh_token" json:"refresh_token,omitempty"`     //
	Scope         *string           `query:"scope" json:"scope,omitempty"`                     //
	TokenType     *consts.TokenType `query:"token_type" json:"token_type,omitempty"`           //
	CodeChallenge *string           `query:"code_challenge" json:"code_challenge,omitempty"`   //
	Code          *string           `query:"code" json:"code,omitempty"`                       //
	CodeExpireAt  *time.Time        `query:"code_expire_at" json:"code_expire_at,omitempty"`   //
	Used          *bool             `query:"used" json:"used,omitempty"`                       //
}

type TokenItem struct {
	ID            uint64           `json:"id,omitempty"`             //
	CreatedAt     time.Time        `json:"created_at,omitempty"`     //
	UserID        uint64           `json:"user_id,omitempty"`        //
	AccessToken   string           `json:"access_token,omitempty"`   //
	RefreshToken  string           `json:"refresh_token,omitempty"`  //
	Scope         string           `json:"scope,omitempty"`          //
	TokenType     consts.TokenType `json:"token_type,omitempty"`     //
	CodeChallenge string           `json:"code_challenge,omitempty"` //
	Code          string           `json:"code,omitempty"`           //
	CodeExpireAt  time.Time        `json:"code_expire_at,omitempty"` //
	SessionID     uint64           `json:"session_id"`
	Session       *models.Session  `json:"session,omitempty"`
	ExpireAt      time.Time        `json:"expire_at"`
	Used          bool             `json:"used,omitempty"` //
}
