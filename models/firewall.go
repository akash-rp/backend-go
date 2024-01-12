package models

import "github.com/google/uuid"

type FirewallDB struct {
	SevenG              bool      `db:"sevenG"`
	SevenGDisabled      []string  `db:"sevenG_disabled"`
	ModSecurity         bool      `db:"modsecurity"`
	ModSecurityParanoia int       `db:"modsecurity_paranoiaLevel"`
	ModSecurityAnomaly  int       `db:"modsecurity_anomalyThreshold"`
	ID                  uuid.UUID `db:"id"`
	SiteID              uuid.UUID `db:"siteid"`
}

type UpdateSevenGFirewall struct {
	App     string   `json:"app"`
	User    string   `json:"user"`
	Disable []string `json:"disable" binding:"required"`
	Enabled bool     `json:"enabled"`
}

type UpdateModSecFirewall struct {
	App              string `json:"app"`
	Enabled          bool   `json:"enabled"`
	ParanoiaLevel    int    `json:"paranoiaLevel"`
	AnomalyThreshold int    `json:"anomalyThreshold"`
}
