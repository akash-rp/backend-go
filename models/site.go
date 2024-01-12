package models

import "github.com/google/uuid"

type AddSite struct {
	AppName       string         `json:"appName" binding:"required"`
	URL           string         `json:"url" binding:"required"`
	UserName      string         `json:"userName" binding:"required"`
	Title         string         `json:"title" binding:"required"`
	AdminUser     string         `json:"adminUser" binding:"required"`
	AdminEmail    string         `json:"adminEmail" binding:"required"`
	AdminPassword string         `json:"adminPassword" binding:"required"`
	Php           string         `json:"php" binding:"required"`
	Domain        ModifiedDomain `json:"domain" binding:"required"`
}

type SiteDetails struct {
	ID             uuid.UUID     `json:"id"`
	ServerID       uuid.UUID     `json:"serverId"`
	Name           string        `json:"name"`
	PHP            string        `json:"php"`
	Staging        uuid.NullUUID `json:"staging,omitempty"`
	Type           int           `json:"type"`
	Authentication bool          `json:"authentication"`
	User           string        `json:"user"`
	UserID         uuid.UUID     `json:"-"`
	Domains        []Domain      `json:"domains"`
	IP             string        `json:"ip"`
}
