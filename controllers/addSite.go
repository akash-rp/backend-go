package controllers

import (
	"backend/db"
	"backend/models"
	"backend/server"
	"backend/site"
	"backend/utilites"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AddSite(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	siteDetails := new(models.AddSite)

	if err := ctx.ShouldBind(&siteDetails); err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	serverDetails, err := server.GetServerDetails(ctx.Param("id"), userId.(string))
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	exists, err := site.CheckIfSiteWithNameExists(serverDetails, siteDetails.AppName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if exists {
		utilites.AbortWithErrorMessage(ctx, "App name exists in the server")
		return
	}

	parsedDomain, err := site.ModifyDomain(siteDetails.URL)
	if err != nil {
		utilites.AbortWithErrorMessage(ctx, "Invalid domain")
		return
	}

	exists, err = site.CheckIfDomainExists(serverDetails.Id, parsedDomain.Url)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if exists {
		utilites.AbortWithErrorMessage(ctx, "Domain name exists in the server")
		return
	}

	//Show routing none only for non-primary domains
	if parsedDomain.Routing == "none" {
		parsedDomain.Routing = "root"
	}

	siteDetails.Domain = parsedDomain
	siteDetailsJSON, _ := json.Marshal(&siteDetails)

	resp, err := http.Post(fmt.Sprintf("http://%s:8081/wp/add", serverDetails.IP), "application/json", bytes.NewReader(siteDetailsJSON))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := new(models.Error)
		json.NewDecoder(resp.Body).Decode(&errorMessage)
		ctx.JSON(resp.StatusCode, errorMessage)
		return
	}
	//fmt.Print(parsedDomain)
	id, _ := uuid.NewRandom()

	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	_, err = tx.Exec(ctx, "INSERT INTO sites(id, serverid, name, php, type, authentication, userid, \"user\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", id, serverDetails.Id, siteDetails.AppName, siteDetails.Php, 1, false, userId, siteDetails.UserName)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	_, err = tx.Exec(ctx, "INSERT INTO domains(url, type, ssl, wildcard, subdomain, routing, siteid, serverid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);", parsedDomain.Url, 1, false, false, parsedDomain.IsSubDomain, parsedDomain.Routing, id, serverDetails.Id)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}

	err = site.InsertFirewallQuery(tx, id)
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

	SitesByServer(ctx)
}
