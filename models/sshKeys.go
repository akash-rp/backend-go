package models

type SshKey struct {
	Username  string `json:"username"`
	Key       string `json:"key"`
	Label     string `json:"label"`
	Timestamp int64  `json:"timestamp"`
}
