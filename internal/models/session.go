package models

type Session struct {
	ID        uint64 `json:"id,omitempty" db:"user_id"`
	AuthToken string `json:"auth_token,omitempty" db:"auth_token"`
}
