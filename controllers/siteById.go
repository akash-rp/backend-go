package controllers

import (
	"backend/db"
	"backend/models"
	"backend/site"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

func GetSiteById(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	siteId := ctx.Param("id")

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

	domainsQuery, _ := tx.Query(ctx, "SELECT id, url, type, ssl, wildcard, subdomain, routing, siteid from domains WHERE siteid = $1", siteId)
	if err != nil {
		fmt.Printf("%+v", err)
		ctx.AbortWithStatus(400)
		return
	}

	domains, err := pgx.CollectRows(domainsQuery, pgx.RowToStructByNameLax[models.Domain])
	if err != nil {
		fmt.Printf("%+v", err)
		ctx.AbortWithStatus(400)
		return
	}

	wp.Domains = domains

	ctx.JSON(http.StatusOK, wp)
}
