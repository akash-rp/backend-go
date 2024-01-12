package controllers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func FirewallBySiteId(ctx *gin.Context) {
	siteId := ctx.Param("id")

	firewallQuery, err := db.DbConn.Query(ctx, "SELECT * FROM firewall WHERE siteid = $1", siteId)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	firewall, err := pgx.CollectOneRow(firewallQuery, pgx.RowToStructByNameLax[models.FirewallDB])
	if err != nil {
		fmt.Printf("%+v", err)
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(200, gin.H{
		"sevenG": gin.H{
			"enabled": firewall.SevenG,
			"disable": firewall.SevenGDisabled,
		},
		"modsecurity": gin.H{
			"enabled":          firewall.ModSecurity,
			"paranoiaLevel":    firewall.ModSecurityParanoia,
			"anomalyThreshold": firewall.ModSecurityAnomaly,
		},
	})
}
