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
	"time"
)

func AddDomain(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

	addDomain := new(models.AddDomain)
	if err := ctx.ShouldBind(&addDomain); err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer tx.Commit(ctx)

	wp, err := site.GetSiteDetails(siteId, userId.(string), tx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	parsedDomain, err := site.ModifyDomain(addDomain.Url)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	exists, err := site.CheckIfDomainExists(wp.ServerID, parsedDomain.Url)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	if exists {
		utilites.AbortWithErrorMessage(ctx, "Domain already exists in server")
		return
	}

	domainJSON, _ := json.Marshal(gin.H{
		"domain": parsedDomain,
		"site":   wp.Name,
	})

	start := time.Now()
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/domain/add", wp.IP), "application/json", bytes.NewReader(domainJSON))
	if err != nil {
		fmt.Print(err)
		ctx.AbortWithStatus(400)
		return
	}
	fmt.Print("time taken", time.Since(start))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := new(models.Error)
		json.NewDecoder(resp.Body).Decode(&errorMessage)
		ctx.JSON(resp.StatusCode, errorMessage)
		return
	}

	result, err := site.InsertDomainQuery(tx, &parsedDomain, addDomain.Type, false, false, wp.ServerID, siteId)
	if err != nil {
		fmt.Print(err)
		ctx.AbortWithStatus(400)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Print(err)
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(200, result)
}
