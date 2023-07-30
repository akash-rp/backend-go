package models

type SystemUsers []struct {
	User  string `json:"user"`
	Sites int    `json:"sites"`
}
