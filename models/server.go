package models

import "github.com/google/uuid"

type ServerList struct {
	Id   uuid.UUID `json:"serverId"`
	Name string    `json:"name"`
	IP   string    `json:"ip"`
}

type ServerDetails struct {
	Id    uuid.UUID   `json:"serverId"`
	Name  string      `json:"name"`
	IP    string      `json:"ip"`
	Stats ServerStats `json:"stats"`
}

type ServerStats struct {
	Cores       string `json:"cores"`
	CPU         string `json:"cpu"`
	TotalMemory string `json:"totalMemory"`
	UsedMemory  string `json:"usedMemory"`
	TotalDisk   string `json:"totalDisk"`
	UsedDisk    string `json:"usedDisk"`
	Bandwidth   string `json:"bandwidth"`
	Os          string `json:"os"`
	Uptime      string `json:"uptime"`
	Loadavg     string `json:"loadavg"`
	Cpuideal    string `json:"cpuideal"`
}

type ServerHealth struct {
	Memory []struct {
		Time  int64 `json:"time"`
		Value int   `json:"value"`
	} `json:"memory"`
	CPU []struct {
		Time  int64 `json:"time"`
		Value int   `json:"value"`
	} `json:"cpu"`
	Load []struct {
		Time  int64 `json:"time"`
		Value int   `json:"value"`
	} `json:"load"`
	Disk []struct {
		Time  int64   `json:"time"`
		Value float64 `json:"value"`
	} `json:"disk"`
}

type ServiceStatus []struct {
	Service string `json:"service"`
	Running bool   `json:"running"`
	Process string `json:"process"`
}
