package models

import "github.com/google/uuid"

type ModifiedDomain struct {
	Url         string `json:"url"`
	IsSubDomain bool   `json:"isSubDomain"`
	Routing     string `json:"routing"`
}

type Domain struct {
	Id        uuid.UUID `json:"id"`
	Url       string    `json:"url"`
	Subdomain bool      `json:"subDomain"`
	Routing   string    `json:"routing"`
	Type      int       `json:"type"`
	Wildcard  bool      `json:"wildcard"`
	Siteid    uuid.UUID `json:"siteId"`
	Ssl       bool      `json:"ssl"`
}

type AddDomain struct {
	Url  string `json:"url" binding:"required"`
	Type int    `json:"type" binding:"required"`
}

type DeleteDomain struct {
	Url string `json:"url" binding:"required"`
}

type UpdateWildcard struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	Wildcard bool      `json:"wildcard"`
	Type     int       `json:"type"`
}

type ChangePrimary struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
