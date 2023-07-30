package models

type SshUsers []struct {
	User    string `json:"user"`
	ID      string `json:"id"`
	IP      string `json:"ip"`
	Login   string `json:"login"`
	Ideal   string `json:"ideal"`
	Process string `json:"process"`
}
