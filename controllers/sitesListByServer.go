package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

func SitesByServer(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	serverId := ctx.Param("id")

	tx, err := db.DbConn.Begin(ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.AbortWithStatus(400)
		return
	}
	defer tx.Commit(ctx)

	sitesQuery, err := tx.Query(ctx, "SELECT * from sites WHERE serverid = $1 AND userid = $2", serverId, userId)
	if err != nil {
		fmt.Printf("%+v", err)
		ctx.AbortWithStatus(400)
		return
	}

	sites, err := pgx.CollectRows(sitesQuery, pgx.RowToStructByNameLax[models.SiteDetails])
	if err != nil {
		fmt.Printf("%+v", err)
		ctx.AbortWithStatus(400)
		return
	}

	domainsQuery, _ := tx.Query(ctx, "SELECT id, url, type, ssl, wildcard, subdomain, routing, siteid from domains WHERE serverid = $1 and type = $2", serverId, 1)
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

	for _, domain := range domains {
		for i, site := range sites {
			if site.ID == domain.Siteid {
				sites[i].Domains = append(sites[i].Domains, domain)
			}
		}
	}
	fmt.Print(sites)
	ctx.JSON(http.StatusOK, sites)
}
