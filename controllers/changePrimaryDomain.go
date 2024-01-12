package controllers

import (
	"backend/db"
	"backend/models"
	"backend/site"
	"backend/utilites"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangePrimaryDomain(ctx *gin.Context) {
	//userId, _ := ctx.Get("userId")
	wpJSON, _ := ctx.Get("site")
	wp := new(models.SiteDetails)
	json.Unmarshal(wpJSON.([]byte), wp)
	siteId := ctx.Param("id")

	newPrimaryDomainId := new(models.ChangePrimary)
	if err := ctx.ShouldBind(&newPrimaryDomainId); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	domain, err := site.GetDomainById(tx, newPrimaryDomainId.Id.String(), siteId)
	if err != nil {
		fmt.Print(err.Error())
		utilites.AbortWithErrorMessage(ctx, "Domain not found.")
		return
	}

	primaryDomain, err := site.GetPrimaryDomain(tx, siteId)
	if err != nil {
		fmt.Print(err.Error())
		utilites.AbortWithErrorMessage(ctx, "Primary domain not found.")
		return
	}

	if domain.Url == primaryDomain.Url {
		utilites.AbortWithErrorMessage(ctx, "Invalid domain type.")
		return
	}

	domainJSON, _ := json.Marshal(gin.H{
		"name":           wp.Name,
		"user":           wp.User,
		"currentPrimary": primaryDomain.Url,
		"newPrimary":     domain.Url,
	})

	resp, err := http.Post(fmt.Sprintf("http://%s:8081/domain/primary", wp.IP), "application/json", bytes.NewReader(domainJSON))
	if err != nil {
		fmt.Print(err)
		ctx.AbortWithStatus(400)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := new(models.Error)
		json.NewDecoder(resp.Body).Decode(&errorMessage)
		ctx.JSON(resp.StatusCode, errorMessage)
		return
	}

	domain.Type = 1
	updatedPrimaryDomain, err := site.UpdateDomainQuery(tx, domain, domain.Id, wp.ID)
	if err != nil {
		fmt.Print(err.Error())
		utilites.AbortWithErrorMessage(ctx, "Failed to update new primary domain type")
		return
	}

	primaryDomain.Type = 2
	aliasDomain, err := site.UpdateDomainQuery(tx, primaryDomain, primaryDomain.Id, wp.ID)
	if err != nil {
		fmt.Print(err.Error())
		utilites.AbortWithErrorMessage(ctx, "Failed to update primary domain type to alias")
		return
	}

	ctx.JSON(200, gin.H{
		"domains": []models.Domain{updatedPrimaryDomain, aliasDomain},
	})
}
