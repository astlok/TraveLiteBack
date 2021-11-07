package models

type Region struct {
	ID uint64 `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
}
