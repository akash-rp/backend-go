package controllers

import (
	"backend/db"
	"backend/models"
	"backend/site"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func DomainWildcard(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

	updateWildcard := new(models.UpdateWildcard)
	if err := ctx.ShouldBind(&updateWildcard); err != nil {
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

	wp, err := site.GetSiteDetails(siteId, userId.(string), tx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	domain, err := site.GetDomainById(tx, updateWildcard.Id.String(), siteId)
	fmt.Print(updateWildcard)
	domain.Wildcard = updateWildcard.Wildcard

	domainJSON, _ := json.Marshal(gin.H{
		"domain": domain,
		"site":   wp.Name,
	})

	resp, err := http.Post(fmt.Sprintf("http://%s:8081/domain/wildcard/update", wp.IP), "application/json", bytes.NewReader(domainJSON))
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

	result, err := site.UpdateDomainQuery(tx, domain, updateWildcard.Id, uuid.MustParse(siteId))
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
