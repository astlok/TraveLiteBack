package models

type User struct {
	ID        uint64 `json:"id,omitempty" db:"id"`
	Email     string `json:"email,omitempty" db:"email"`
	Nickname  string `json:"nickname,omitempty" db:"nickname"`
	Password  string `json:"password,omitempty" db:"password"`
	IMG       string `json:"img,omitempty" db:"img"`
	AuthToken string `json:"auth_token,omitempty" db:"auth_token"`
}
