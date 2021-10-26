package models

type Session struct {
	UserID uint64 `json:"user_id,omitempty" db:"user_id"`
	AuthToken string `json:"auth_token,omitempty" db:"auth_token"`
}
