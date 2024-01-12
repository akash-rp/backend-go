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

func DeleteDomain(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")
	domainId := ctx.Param("domainId")

	if domainId == "" {
		fmt.Print("no domainId")

		utilites.AbortWithErrorMessage(ctx, "Domain Id is required.")
		return
	}
	fmt.Print("Stage 1")

	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print("Something went wrong")
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer tx.Commit(ctx)
	fmt.Print("Stage 2")

	wp, err := site.GetSiteDetails(siteId, userId.(string), tx)
	if err != nil {
		fmt.Print("Site not found")
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	fmt.Print("Stage 3")

	parsedDomain, err := site.GetDomainById(tx, domainId, wp.ID.String())
	if err != nil {
		fmt.Print("Domain not found")
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	fmt.Print("Stage 4")

	domainJSON, _ := json.Marshal(gin.H{
		"domain": parsedDomain.Url,
		"site":   wp.Name,
	})
	fmt.Print(string(domainJSON))
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/domain/delete", wp.IP), "application/json", bytes.NewReader(domainJSON))
	if err != nil {
		fmt.Print("Failed to connect to agent")
		fmt.Print(err)
		ctx.AbortWithStatus(400)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := new(models.Error)
		json.NewDecoder(resp.Body).Decode(&errorMessage)
		fmt.Print("Stage 5")
		ctx.JSON(resp.StatusCode, errorMessage)
		return
	}

	err = site.DeleteDomainQuery(tx, parsedDomain.Id.String(), siteId)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(200, gin.H{})
}
