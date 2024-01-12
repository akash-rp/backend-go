package models

type SystemUsers []struct {
	User  string `json:"user"`
	Sites int    `json:"site"`
}

type SystemUserCred struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type DeleteSystemUser struct {
	User string `json:"user"`
}
