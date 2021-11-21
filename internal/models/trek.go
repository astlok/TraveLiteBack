package models

type Trek struct {
	ID   uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Difficult   int      `json:"difficult" db:"difficult"`
	Days        int      `json:"days" db:"days"`
	Things      []string `json:"things"`
	Description string   `json:"description" db:"description"`
	File        string   `json:"file" db:"file"`
	Region      string   `json:"region" db:"region_name"`
	Rating      float64 `json:"rating" db:"rating"`
}
